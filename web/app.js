var app = angular.module('CocktailApp', ['ngRoute']);

app.config(function($locationProvider) {
    $locationProvider.hashPrefix('');
});

app.config(function($routeProvider) {
    $routeProvider
        .when('/', {
            templateUrl: 'menu.html'
        })
        .when('/cocktails', {
            templateUrl: 'cocktail-list.html',
            controller: 'CocktailListController'
        })
        .when('/cocktails/:id', {
            templateUrl: 'cocktail-detail.html',
            controller: 'CocktailDetailController'
        })
        .when('/ingredients', {
            templateUrl: 'ingredient-list.html',
            controller: 'IngredientListController'
        })
        .when('/ingredients/:id', {
            templateUrl: 'ingredient-detail.html',
            controller: 'IngredientDetailController'
        })
        .when('/search', {
            templateUrl: 'search.html',
            controller: 'SearchController'
        })
        .when('/user/fridge', {
            templateUrl: 'fridge.html',
            controller: 'FridgeController'
        })
        .when('/user/possible', {
            templateUrl: 'possible.html',
            controller: 'PossibleController'
        })
        .when('/user/shopping', {
            templateUrl: 'shopping.html',
            controller: 'ShoppingController'
        })
        // Eigene Ansicht für den Live-Feed, die parallel zur Detailseite geöffnet werden kann.
        .when('/feed', {
            templateUrl: 'rating-feed.html',
            controller: 'RatingFeedController'
        })
        .otherwise({
            redirectTo: '/'
        });
});

app.controller('CocktailListController', function($scope, $http) {
    $http.get('/api/cocktails')
        .then(function(response) {
            $scope.cocktails = response.data;
        });
});

app.controller('CocktailDetailController', function($scope, $routeParams, $http) {
    $http.get('/api/cocktails/' + $routeParams.id)
        .then(function(response) {
            $scope.cocktail = response.data;
        });

    // Eine Bewertung abschicken und das Formular erst nach Erfolg zurücksetzen.
    $scope.rating = {};
    $scope.submitRating = function() {
        $scope.savingRating = true;
        $scope.ratingMessage = '';
        $http.post('/api/cocktails/' + $routeParams.id + '/ratings', $scope.rating)
            .then(function() {
                $scope.rating = {};
                $scope.ratingMessage = 'Bewertung gespeichert.';
            }, function(response) {
                $scope.ratingMessage = (response.data && response.data.error) ||
                    'Bewertung konnte nicht gespeichert werden.';
            })
            .finally(function() {
                $scope.savingRating = false;
            });
    };
});

app.controller('IngredientListController', function($scope, $http) {
    $http.get('/api/ingredients')
        .then(function(response) {
            $scope.ingredients = response.data;
        });
});

app.controller('IngredientDetailController', function($scope, $routeParams, $http) {
    $http.get('/api/ingredients/' + $routeParams.id)
        .then(function(response) {
            $scope.ingredient = response.data;
        });
});

app.controller('SearchController', function($scope, $http, $location) {
    $scope.query = $location.search().query || '';

    if ($scope.query) {
        // Führe die Suche durch
        $http.get('/api/search', {params: {query: $scope.query}})
            .then(function (response) {
                $scope.results = response.data;
            })
    }

    $scope.search = function() {
        // Aktualisiere die URL mit dem Such-Query
        $location.search('query', $scope.query);
    };
});

app.controller('FridgeController', function($scope, $http) {
    $scope.updateFridge = function() {
        $scope.inFridge = $scope.fridge.filter(function(ingredient) {
            return ingredient.inFridge;
        });
        $scope.notInFridge = $scope.fridge.filter(function(ingredient) {
            return !ingredient.inFridge;
        });
    }

    // Hilfsfunktion zum Ändern einer Zutat
    $scope.updateIngredient = function(id, inFridge) {
        $scope.fridge.forEach(function(ingredient) {
            if (ingredient.id === id) {
                ingredient.inFridge = inFridge;
            }
        });
        $scope.updateFridge();
        $http.patch('/api/user/fridge/' + id, { inFridge: inFridge })
    }

    // Initialisierung beim Laden der Seite
    $http.get('/api/user/fridge')
        .then(function(response) {
            $scope.fridge = response.data;
            $scope.updateFridge();
        });

    // Funktion zum Hinzufügen einer Zutat zum Kühlschrank
    $scope.addToFridge = function(id) {
        // Aktualisiere die inFridge-Eigenschaft der Zutat
        $scope.updateIngredient(id, true)
    };

    // Funktion zum Entfernen einer Zutat aus dem Kühlschrank
    $scope.removeFromFridge = function(id) {
        $scope.updateIngredient(id, false)
    };
});

app.controller('PossibleController', function($scope, $http) {
    $http.get('/api/user/possible')
        .then(function(response) {
            $scope.cocktails = response.data;
        });
});

app.controller('ShoppingController', function($scope, $http) {
    $http.get('/api/user/shopping')
        .then(function(response) {
            $scope.ingredients = response.data;
        });
});


// Der Feed kombiniert die gespeicherte Historie mit neuen Einträgen aus MongoDB.
app.controller('RatingFeedController', function($scope, $http) {
    $scope.ratings = [];
    $scope.feedStatus = 'Live-Verbindung wird aufgebaut …';

    // Einträge können gleichzeitig über die Historie und den Stream ankommen.
    function mergeRatings(incoming) {
        var byId = new Map();
        $scope.ratings.concat(incoming).forEach(function(rating) {
            byId.set(rating.id, rating);
        });
        $scope.ratings = Array.from(byId.values()).sort(function(a, b) {
            return Date.parse(b.createdAt) - Date.parse(a.createdAt) ||
                b.id.localeCompare(a.id);
        }).slice(0, 20);
    }

    var stream = new EventSource('/api/ratings/stream');

    function loadHistory() {
        $http.get('/api/ratings').then(function(response) {
            mergeRatings(response.data);
            $scope.historyError = '';
        }, function() {
            $scope.historyError = 'Bewertungen konnten nicht geladen werden.';
        });
    }

    // Nach dem Start und jeder Wiederverbindung schließt die Historie mögliche Lücken.
    stream.addEventListener('ready', function() {
        loadHistory();
        $scope.$evalAsync(function() {
            $scope.feedStatus = 'Live verbunden';
        });
    });
    stream.addEventListener('rating', function(event) {
        var rating = JSON.parse(event.data);
        // EventSource läuft außerhalb von AngularJS.
        $scope.$evalAsync(function() {
            mergeRatings([rating]);
        });
    });
    stream.onerror = function() {
        $scope.$evalAsync(function() {
            $scope.feedStatus = 'Live-Verbindung unterbrochen.';
        });
        // Die Historie bleibt auch bei einem ausgefallenen Stream zugänglich.
        loadHistory();
    };

    // Beim Wechsel der Route soll der Browser die Verbindung schließen.
    $scope.$on('$destroy', function() {
        stream.close();
    });
});
