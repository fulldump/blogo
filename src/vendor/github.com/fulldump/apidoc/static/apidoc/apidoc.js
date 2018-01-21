function is_param(p) {
	return p.length > 1 && p[0] === '{'
}

function normalize_header(h) {
	return h.split('-').map(function(w) {
		return w[0].toUpperCase() + w.substr(1).toLowerCase();
	}).join('-');
}

angular

.module('ApiDoc', [])

.directive("contenteditable", function() {
  return {
    restrict: "A",
    require: "ngModel",
    link: function(scope, element, attrs, ngModel) {

      function read() {
        ngModel.$setViewValue(element.html());
      }

      ngModel.$render = function() {
        element.html(ngModel.$viewValue || "");
      };

      element.bind("blur keyup change", function() {
        scope.$apply(read);
      });
    }
  };
})

.controller('MainCtrl', function($http) {

	this.doc = {};

	$http({
		method: 'GET',
		url: 'json',
	}).then(
		function ok(response) {
			that.doc = response.data;

			var endpoints = that.doc.endpoints;
			for (var url in endpoints) {
				var endpoint = endpoints[url];

				endpoint.parameters = {};

				endpoint.url = [];
				var parts = url.split("/");
				parts.shift();
				for (var i in parts) {
					var part = parts[i];
					endpoint.url.push({
						name: part,
						isparam: is_param(part), 
					});
					if (is_param(part)) {
						endpoint.parameters[part] = {
							name: part,
							value: part,
						};
					}
				}

				var methods = endpoint.methods;
				for (var i in methods) {
					var method = methods[i];
					if ('POST' == i || 'PATCH' == i || 'PUT' == i) {
						method.hasbody = true;
						method.bodyshow = false;
						method.body = '{\n    \n}';
					} else {
						method.hasbody = false;
					}
				}

			}
		},
		function error(response) {

		}
	);

	this.valid_json = function(s) {
		try {
			JSON.parse(s);
			return true;
		} catch (e) {}
		return false;
	};

	this.request = function(method, endpoint, item){
		if (endpoint.methods[method].hasbody && !endpoint.methods[method].bodyshow) {
			endpoint.methods[method].bodyshow = true;
			return;
		}

		var url = [];
		for (var i in endpoint.url) {
			var part = endpoint.url[i].name;
			if (is_param(part)) {
				var value = endpoint.parameters[part].value;
				url.push(encodeURIComponent(value));
			} else {
				url.push(part);
			}
		}

		var cb = function ok(response) {
			item.response = 'HTTP/1.1 ' + response.status + ' ' + response.statusText + '\n';

			var headers = response.headers();
			for (var h in headers) {
				item.response += normalize_header(h) + ': ' + headers[h] + '\n';
			}
			item.response += '\n';

			try {
				var json = JSON.parse(response.data);
				item.response += JSON.stringify(json, null, 4);
			} catch (e) {
				item.response += response.data;
			}
		}

		$http({
			method: method,
			url: '/'+url.join('/'),
			transformResponse: [function (data) {
				return data;
  			}],
  			data: endpoint.methods[method].body,
		}).then(cb, cb);
	};

	var that = this;

})

.directive('marked', function() {
	return {
		restrict: 'E',
		scope: {
			content: '=',
		},
		link: function(scope, element, attrs) {
			var div = element[0];
			div.innerHTML = marked(scope.content);
			// TODO: colorize code
		},
	};
})

