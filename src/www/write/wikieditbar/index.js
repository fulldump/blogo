(function(context) {

    var Button = function(name) {

        this.name = name;

        this.dom = document.createElement('button');
        this.dom.setAttribute('id', name);
    };

    var Group = function(name) {

        this.name = name;

        this.dom = document.createElement('div');
        this.dom.classList.add('group');
        this.dom.setAttribute('id', name);

        this.buttons = {};

        this.slide = null;
    };

    Group.prototype.addButton = function(button_name, on_click) {
        if (!(button_name in this.buttons)) {
            var button = new Button(button_name);
            this.buttons[button_name] = button;
            this.dom.appendChild(button.dom);

            if (on_click) {
                button.dom.addEventListener('click', on_click, true);
            }
        }

        return this.buttons[button_name];
    };

    Group.prototype.addSlideButton = function(button_name, on_click) {
        var button = this.addButton(button_name, on_click);

        if (null === this.slide) {
            this.slide = document.createElement('div');
            this.dom.appendChild(this.slide);
        }
        this.slide.appendChild(button.dom);
    };

    Group.prototype.addButtons = function(button_name, on_click) {
        var button = this.addButton(button_name, on_click);

        return this;
    };

    Group.prototype.addSlideButtons = function(button_name, on_click) {
        var button = this.addSlideButton(button_name, on_click);

        return this;
    };


    var WikiEditBar = function() {

        var dom = this.dom = document.createElement('div');
        this.dom.dataset.component = 'WikiEditBar';

        this.bar = document.createElement('div');
        this.bar.setAttribute('id', 'bar');
        this.dom.appendChild(this.bar);

        this.content = document.createElement('div');
        this.content.classList.add('max-width');
        this.bar.appendChild(this.content);

        this.spacer = document.createElement('div');
        this.spacer.setAttribute('id', 'spacer');
        this.dom.appendChild(this.spacer);

        this.groups = {};

        window.addEventListener('scroll', function(e) {
            //var scrollTop = document.body.scrollTop;  // Works in chrome
            var scrollTop = window.scrollY; // Works in chrome and firefox
            if (scrollTop >= dom.offsetTop) {
                dom.classList.add('fixed');
            } else {
                dom.classList.remove('fixed');
            }

            if (scrollTop >= dom.offsetTop - dom.clientHeight) {
                dom.classList.add('shadowed');
            } else {
                dom.classList.remove('shadowed');
            }
        }, true);

    };

    WikiEditBar.prototype.attach = function(selector, callback) {
        this.dom.querySelector(selector).addEventListener('click', callback, true);
    };

    WikiEditBar.prototype.addGroup = function(group_name) {
        if (!(group_name in this.groups)) {
            var group = new Group(group_name);
            this.groups[group_name] = group;
            this.content.appendChild(group.dom);
        }

        return this.groups[group_name];
    };

    WikiEditBar.prototype.addButton = function(group_name, button_name) {
        if (!(group_name in this.groups)) {
            return;
        }

        var group = this.groups[group_name];
        if (button_name in group.buttons) {
            return;
        }

        var dom = document.createElement('button');
        dom.setAttribute('id', button_name);
        group.dom.appendChild(dom);

        group.buttons[button_name] = {
            dom: dom,
        };
    };


    window.WikiEditBar = WikiEditBar;


})(window);
