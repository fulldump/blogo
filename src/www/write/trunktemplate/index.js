
(function(namespace){

    function execAll(regex, text) {
        var result = {};

        while ((match = regex.exec(text)) !== null) {
            result[match[1]] = match.index;
        }
        return result;
    }

    var TrunkTemplate = function(dom) {
        this.items = {};
        this.inspect(dom);
    }

    TrunkTemplate.prototype.pattern = /{{([^}]+)}}/mg;

    TrunkTemplate.prototype.evaluate = function(type, text, object) {
        var matches = execAll(this.pattern, text);

        for (var k in matches) {
            if (!(k in this.items)) {
                this.items[k] = {
                    'places': [],
                    'value': '',
                };
            }
            this.items[k].places.push({
                type: type,
                object: object,
                text: text
            });
        }
    };

    TrunkTemplate.prototype.inspect = function(dom) {
        // Inspect attributes
        var attributes = dom.attributes;
        for (var i=0; i<attributes.length; i++) {
            var attribute = attributes[i];
            this.evaluate('attribute', attribute.value, attribute);
        }

        var childNodes = dom.childNodes;
        for (var i=0; i<childNodes.length; i++) {
            var childNode = childNodes[i];
            var nodeType = childNode.nodeType;
            if (1 == nodeType) {
                this.inspect(childNodes[i]);
            } else if (3 == nodeType) {
                this.evaluate('text', childNode.textContent, childNode);
            }
        }

    };

    TrunkTemplate.prototype.set = function(key, value) {
        if (key in this.items) {
            var item = this.items[key];
            item.value = value;
            var places = item.places;
            for (var i in places) {
                var place = places[i];
                // Replace all matches
                var matches = execAll(this.pattern, place.text);
                var text = place.text;
                for (var word in matches) {
                    text = text.replace('{{'+word+'}}', this.items[word].value);
                }
                if ('text' == place.type) {
                    place.object.textContent = text;
                } else if ('attribute' == place.type) {
                    place.object.value = text;
                }
            }

        }
        return value;
    };



    namespace.TrunkTemplate = TrunkTemplate;

})(window);
