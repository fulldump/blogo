/**
 * Terrain
 * Text editor for HTML5 browsers.
 *
 * Typical example:
 *
 *	var terrain = new Terrain(document.getElementById('my-document'));
 *	terrain.enableEditor();
 *
 */

(function(context) {

    'use strict';

    var debug_enabled = false;

    var headings = ['H2', 'H3', 'H4', 'H5', 'H6'];

    function debug(item) {
        if (debug_enabled) {
            console.log(item);
        }
    };

    var Terrain = function(dom) {
        this.dom = dom;

        this._last_selection = null;
        this._parts = [];
    };

    Terrain.prototype.getSelection = function() {
        var selection = document.getSelection();

        // Check if anchorNode is inside terrain
        var path = [];
        var element = selection.anchorNode;
        while (element != this.dom && element !== null) {
            path.unshift(element);
            element = element.parentNode;
        }

        if (element === null) {
            return this._last_selection;
        }

        this._last_selection = {
            isCollapsed: selection.isCollapsed,
            anchorNode: selection.anchorNode,
            anchorOffset: selection.anchorOffset,
            focusNode: selection.focusNode,
            focusOffset: selection.focusOffset,
            path: path,
        }

        return this._last_selection;
    };

    Terrain.prototype.enableEditor = function() {
        this.dom.setAttribute('component', 'Terrain');
        this.dom.setAttribute('contentEditable', true);
        this.dom.focus();
        this.checkEmpty();

        var that = this;

        this.dom.addEventListener('blur', function(event) {
            that.clean();
        }, true);

        this.dom.addEventListener('keydown', function(event) {
            that.processEvent(event);
        }, true);

        this.dom.addEventListener('click', function(event) {
            that.getSelection();
        }, true);

        this.dom.addEventListener('paste', function(event) {
            that.processEventPaste(event);
        }, true);

    };

    Terrain.prototype.processEventPaste = function(event) {
        var path = this.getDomPath();
        if (path.length > 0 && path[0].nodeName == 'CODE') {
            var node = path[0];
            event.preventDefault();
            if (true || event.clipboardData.types.indexOf('text/plain') > -1) {
                var text = event.clipboardData.getData('text/plain');
                var selection = window.getSelection();
                var anchor = selection.anchorNode;
                var offset = selection.anchorOffset;

                // Split current text node (if proceed)
                if (3 == anchor.nodeType) {
                    var pre = document.createTextNode(anchor.nodeValue.substring(0, offset));
                    anchor.parentNode.insertBefore(pre, anchor);
                    anchor.nodeValue = anchor.nodeValue.substring(offset);
                }

                var lines = text.split('\n');
                var node_reference = null;
                if (1 == anchor.nodeType) {
                    node_reference = anchor.childNodes[offset];
                } else if (3 == anchor.nodeType) {
                    node_reference = anchor;
                }
                for (var i in lines) {
                    var t = document.createTextNode(lines[i].replace(/ /mg, "\u00a0"));
                    if (1 == anchor.nodeType) {
                        anchor.insertBefore(t, node_reference);
                    } else if (3 == anchor.nodeType) {
                        anchor.parentNode.insertBefore(t, anchor);
                    }

                    if (i < lines.length - 1) {
                        var br = document.createElement('br');
                        if (1 == anchor.nodeType) {
                            anchor.insertBefore(br, node_reference);
                        } else if (3 == anchor.nodeType) {
                            anchor.parentNode.insertBefore(br, anchor);
                        }
                    }
                }
                this.putCursorAt(node_reference);
            }
        } else {
            var that = this;
            setTimeout(function() {
                that.clean(this.dom);
            }, 100);
        }
    };

    Terrain.prototype.processEvent = function(event) {
        var path = this.getDomPath();
        var selection = this.getSelection();
        var anchorNode = selection.anchorNode;
        var anchorOffset = selection.anchorOffset;

        if ('keydown' == event.type) {
            this.checkEmpty();

            var node = path[0];

            if (13 == event.keyCode) { // Enter key

                if (path.length > 0 && (path[0].nodeName == 'CODE' || path[0].nodeName == 'BLOCKQUOTE')) {
                    event.stopPropagation();
                    event.preventDefault();

                    if (null === node.lastChild || 'BR' != node.lastChild.tagName) {
                        var br = document.createElement('br');
                        node.appendChild(br);
                    }

                    if (this.isCursorAtBegin(node)) {
                        var p = document.createElement('p');
                        p.innerHTML = '<br>';
                        this.dom.insertBefore(p, node);
                        this.putCursorAt(p);
                        return;
                    }

                    if (
                        // (node.lastChild === anchorNode && 3 === anchorNode.nodeType && ''==anchorNode.nodeValue)  // Chrome
                    // ||
                        (node.lastChild.previousSibling === anchorNode && null !== anchorNode.previousSibling && 1 == anchorNode.previousSibling.nodeType && 'BR' == anchorNode.previousSibling.nodeName && 3 === anchorNode.nodeType && '' == anchorNode.nodeValue) // Gecko
                    ) {
                        var p = document.createElement('p');
                        p.innerHTML = '<br>';
                        this.dom.insertBefore(p, node.nextSibling);
                        this.putCursorAt(p);
                        return;
                    }

                    if (anchorNode.nodeType == 3) {
                        var nodeValue = anchorNode.nodeValue;
                        var pre = document.createTextNode(nodeValue.substring(0, anchorOffset));
                        anchorNode.parentNode.insertBefore(pre, anchorNode);
                        var br = document.createElement('br');
                        anchorNode.parentNode.insertBefore(br, anchorNode);
                        anchorNode.nodeValue = nodeValue.substring(anchorOffset);
                    } else {
                        var node_reference = anchorNode.childNodes[anchorOffset];
                        var jump = document.createElement('br');
                        anchorNode.insertBefore(jump, node_reference);
                    }
                } else if (path.length > 0 && 'FIGURE' == path[0].nodeName) {
                    event.stopPropagation();
                    event.preventDefault();

                    var p = document.createElement('p');
                    p.innerHTML = '<br>';
                    this.dom.insertBefore(p, node.nextSibling);
                    this.putCursorAt(p);
                } else if (headings.indexOf(node.nodeName) >= 0) {
                    event.stopPropagation();
                    event.preventDefault();

                    var p = document.createElement('p');
                    p.innerHTML = '<br>';
                    if (this.isCursorAtBegin(node)) {
                        this.dom.insertBefore(p, node);
                        this.putCursorAt(node);
                    } else {
                        this.dom.insertBefore(p, node.nextSibling);
                        this.putCursorAt(p);
                    }
                }

            } else if (9 == event.keyCode) { // Tab key
                event.preventDefault();
                event.stopPropagation();

                if (path.length > 0) {
                    var path_0 = path[0];
                    if ('CODE' == path_0.nodeName) {
                        document.execCommand('insertText', false, '    ');
                    } else if ('UL' == path_0.nodeName || 'OL' == path_0.nodeName) {
                        if (event.shiftKey) {
                            document.execCommand('outdent', false, null);
                        } else {
                            document.execCommand('indent', false, null);
                        }
                        this.dom.focus();
                    } else if (['H2', 'H3', 'H4', 'H5', 'H6'].indexOf(path_0.nodeName) >= 0) {
                        if (event.shiftKey) {
                            var movement = {
                                H2: 'H2',
                                H3: 'H2',
                                H4: 'H3',
                                H5: 'H4',
                                H6: 'H5'
                            };
                        } else {
                            var movement = {
                                H2: 'H3',
                                H3: 'H4',
                                H4: 'H5',
                                H5: 'H6',
                                H6: 'H6'
                            };
                        }

                        var h = document.createElement(movement[path_0.nodeName]);
                        this.replaceTag(h, path_0);
                        this.putCursorAt(h);
                    }
                }
            }
        } else if ('') {

        }
    };

    Terrain.prototype.highlightCodes = function() {
        var languages = ["xml", "bash", "cpp", "css", "markdown", "haskell", "go", "java", "javascript", "json", "php", "python", "sql"];
        hljs.configure({useBR: true, languages:languages});
        var root = this.dom;
        var n = root.childNodes.length;
        for (var i = 0; i < n; i++) {
            var node = root.childNodes[i];
            if (1 == node.nodeType && 'CODE' == node.nodeName) {
                hljs.highlightBlock(node, null, false);
                var c = node.getAttribute('class').split(' ');
                var lang = document.createElement('span');
                lang.classList.add('lang');
                lang.innerHTML = c[1];
                node.insertBefore(lang, node.firstChild);

                var numlines = node.innerHTML.split('<br>').length;
                var gutter = document.createElement('span');
                gutter.classList.add('gutter');
                var gutter_numbers = [];
                for (var j=1; j<=numlines; j++) {
                    gutter_numbers.push(''+j+'');
                }
                gutter.innerHTML = gutter_numbers.join('<br>');
                node.insertBefore(gutter, node.firstChild);
            }
        }
    };

    Terrain.prototype.removeEmptyParagraphs = function() {
        var not_empty = /\S+/m;
        var root = this.dom;
        var childNodes = root.childNodes;
        var n = childNodes.length;
        for (var i = n-1; i >=0 ; i--) {
            var node = childNodes[i];
            if (1 == node.nodeType && 'P' == node.nodeName && null === not_empty.exec(node.textContent)) {
                root.removeChild(node);
            }
        }
    };

    Terrain.prototype.generateIndex = function() {
        var index = [];

        var processH = function(item, level) {
            index.push({
                item: item,
                level: level
            });
        };

        var root = this.dom;
        var n = root.childNodes.length;
        for (var i = 0; i < n; i++) {
            var node = root.childNodes[i];
            if (1 == node.nodeType) {
                var tagName = node.tagName;
                switch (tagName) {
                    case 'H2':
                        processH(node, 1);
                        break;
                    case 'H3':
                        processH(node, 2);
                        break;
                    case 'H4':
                        processH(node, 3);
                        break;
                    case 'H5':
                        processH(node, 4);
                        break;
                    case 'H6':
                        processH(node, 5);
                        break;
                }
            }
        }

        var current_level = 0;
        var result = '';

        var identation = [];

        var i = 0;
        while (i < index.length) {
            var level = index[i].level;
            if (current_level < level) {
                identation.push(0);
                result += ('<ol>');
                current_level++;
            } else if (current_level > level) {
                identation.pop();
                result += ('</ol>');
                current_level--;
            } else {
                identation.push(1 + identation.pop());

                var item = index[i].item;
                item.id = 'title_' + identation.join('-');
                item.innerHTML = '<a href="#' + item.id + '"><span class="identation">' + identation.join('.') + '</span>' + item.innerHTML + '</a>';
                result += ('<li>' + item.innerHTML + '');

                i++;
            }
        }

        var div = document.createElement('div');
        div.className = 'index';
        div.innerHTML = result;

        this.dom.insertBefore(div, this.dom.firstChild);
    };

    Terrain.prototype.isCursorAtBegin = function(node) {
        var selection = this.getSelection();
        var anchorOffset = selection.anchorOffset;
        var anchorNode = selection.anchorNode;

        if (anchorNode === this.dom) {
            return true;
        }
        if (anchorOffset != 0) {
            return false;
        }
        if (anchorNode.parentNode.childNodes[0] == anchorNode) {
            return true;

            // TODO: fix this
            return this.isCursorAtBegin(anchorNode.parentNode);
        }
        return false;
    };

    Terrain.prototype.isCursorAtEnd = function(node) {
        var selection = this.getSelection();
        var anchorOffset = selection.anchorOffset;
        var anchorNode = selection.anchorNode;

        // TODO: check this recursively until 'node'
        // Is last node?
        var n = anchorNode;
        while (n !== node || null === n) {
            if (n != n.parentNode.lastChild) {
                return false;
            }
            n = n.parentNode;
        }

        if (3 == anchorNode.nodeType) {
            if (anchorOffset == anchorNode.nodeValue.length) {
                return true;
            } else {
                return false;
            }
        } else if (1 === anchorNode.nodeType) {
            return true;
        }
    };

    Terrain.prototype.formatBold = function() {
        document.execCommand('bold', false, null);
        this.dom.focus();
    };

    Terrain.prototype.formatItalic = function() {
        document.execCommand('italic', false, null);
        this.dom.focus();
    };

    Terrain.prototype.formatUnderline = function() {
        document.execCommand('underline', false, null);
        this.dom.focus();
    };

    Terrain.prototype.formatStrike = function() {
        document.execCommand('strikeThrough', false, null);
        this.dom.focus();
    };

    Terrain.prototype.formatBlock = function(block) {
        document.execCommand('formatblock', false, block);
        this.dom.focus();
    }

    Terrain.prototype.putCursorAt = function(node, offset) {
        debug('putCursorAt');
        var range = document.createRange();
        range.setStart(node, offset);

        var selection = window.getSelection();
        selection.removeAllRanges();
        selection.addRange(range);
    };

    Terrain.prototype.removeRootBr = function() {
        debug('removeRootBr');
        var n = this.dom.childNodes.length;
        for (var i = n - 1; i >= 0; i--) {
            var node = this.dom.childNodes[i];
            if (node.nodeType == 1 && node.tagName == 'BR') {
                this.dom.removeChild(this.dom.childNodes[i]);
            }
        }
    };

    Terrain.prototype.getDomPath = function(element) {
        if (!element) {
            return this.getSelection().path;
        }

        var path = [];

        while (element != this.dom) {
            path.unshift(element);
            element = element.parentNode;
        }
        return path;
    };

    Terrain.prototype.insertElement = function(element) {
        var selection = this.getSelection();
        var anchorNode = selection.anchorNode;
        if (3 == anchorNode.nodeType) { // Text node
            var pre = document.createTextNode(anchorNode.nodeValue.substring(0, selection.anchorOffset));
            anchorNode.nodeValue = anchorNode.nodeValue.substring(selection.anchorOffset);
            anchorNode.parentNode.insertBefore(pre, anchorNode);
            anchorNode.parentNode.insertBefore(element, anchorNode);
        } else if (1 == anchorNode.nodeType) { // Tag node
            var node = anchorNode.childNodes[selection.anchorOffset];
            anchorNode.insertBefore(element, node);
        }
    };

    Terrain.prototype.insertCode = function() {
        var selection = this.getSelection();

        if (! selection.collapsed && 3 == selection.anchorNode.nodeType && selection.anchorNode == selection.focusNode) {
            document.execCommand("insertHTML", false, "<code>"+ document.getSelection()+"</code>");
        } else {
            var code = document.createElement('code');
            code.innerHTML = '<br>';

            var reference = this.getDomPath()[0];
            this.dom.insertBefore(code, reference.nextSibling);
            this.putCursorAt(code);
            this.dom.focus();
        }
    };

    Terrain.prototype.insertExternalLink = function(href) {
        var selection = this.getSelection();
        if (! selection.collapsed && 3 == selection.anchorNode.nodeType && selection.anchorNode == selection.focusNode) {
            document.execCommand("insertHTML", false, "<a class=\"external\" href=\""+href+"\">"+ document.getSelection()+"</a>");
        } else {
            document.execCommand("insertHTML", false, "<a class=\"external\" href=\""+href+"\">"+href+"</a>");
        }
    };

    Terrain.prototype.insertInternalLink = function(href) {
        var selection = this.getSelection();
        if (! selection.collapsed && 3 == selection.anchorNode.nodeType && selection.anchorNode == selection.focusNode) {
            document.execCommand("insertHTML", false, "<a class=\"internal\" href=\""+href+"\">"+ document.getSelection()+"</a>");
        } else {
            document.execCommand("insertHTML", false, "<a class=\"internal\" href=\""+href+"\">"+href+"</a>");
        }
    };

    Terrain.prototype.insertFileLink = function(href, name, mime) {
        var a = document.createElement('a');
        a.href = href;
        a.textContent = name;
        a.classList.add('file');
        a.setAttribute('mime', mime);
        this.insertElement(a);
        return a;
    };

    Terrain.prototype.insertQuotes = function() {
        var selection = this.getSelection();

        var code = document.createElement('blockquote');
        code.innerHTML = '<br>';

        var reference = this.getDomPath()[0];
        this.dom.insertBefore(code, reference.nextSibling);
        this.putCursorAt(code);
        this.dom.focus();
    };

    Terrain.prototype.insertUl = function() {
        var ul = document.createElement('ul');
        ul.innerHTML = '<li></li>';

        var reference = this.getDomPath()[0];
        this.dom.insertBefore(ul, reference.nextSibling);
        this.putCursorAt(ul);
        this.dom.focus();
    };

    Terrain.prototype.insertOl = function() {
        var ul = document.createElement('ol');
        ul.innerHTML = '<li></li>';

        var reference = this.getDomPath()[0];
        this.dom.insertBefore(ul, reference.nextSibling);
        this.putCursorAt(ul);
        this.dom.focus();
    };

    Terrain.prototype.insertImage = function(src) {
        var figure = document.createElement('figure');
        var img = document.createElement('img');
        img.src = src;
        figure.appendChild(img);
        var figcaption = document.createElement('figcaption');
        figure.appendChild(figcaption);

        var reference = this.getDomPath()[0];
        this.dom.insertBefore(figure, reference.nextSibling);
        this.putCursorAt(figcaption);
        this.dom.focus();
        return img;
    };

    Terrain.prototype.checkEmpty = function() {
        debug('checkEmpty');
        this.removeRootBr();
        if (this.dom.childNodes.length == 0) {
            var p = document.createElement('p');
            p.innerHTML = '<br>';
            this.dom.appendChild(p);

            this.putCursorAt(p);
        }
    };

    Terrain.prototype.cleanAttributes = function(node, excluded) {
        debug('cleanAttributes');

        var attributes = node.attributes;
        for (var i = attributes.length - 1; i >= 0; i--) {
            var attributeName = attributes[i].name;
            if (!excluded || -1 == excluded.indexOf(attributeName)) {
                node.removeAttribute(attributeName);
            }
        }
    };

    Terrain.prototype.trimText = function(node) {
        debug('trimText');
        if (node.parentNode.childNodes[0] == node) {
            node.nodeValue = node.nodeValue.replace(/^[ \t\n\r]+/gm, '');
        }
    };

    Terrain.prototype.cleanList = function(parent) {
        debug('clean UL|OL');
        this.cleanAttributes(parent);

        var n = parent.childNodes.length;
        for (var i = n - 1; i >= 0; i--) {
            var node = parent.childNodes[i];
            if (1 == node.nodeType) {
                if ('LI' == node.nodeName) {
                    this.cleanText(node);
                } else if ('UL' == node.nodeName || 'OL' == node.nodeName) {
                    this.cleanList(node);
                } else {
                    parent.removeChild(node);
                }
            } else {
                parent.removeChild(node);
            }
        }

        if (0 == parent.childNodes.length) {
            parent.parentNode.removeChild(parent);
        }
    };

    Terrain.prototype.cleanText = function(parent, exclude) {
        if (!exclude) {
            exclude = [];
        }
        var untouched = [];

        debug('cleanText');
        var n = parent.childNodes.length;
        for (var i = n - 1; i >= 0; i--) {
            var node = parent.childNodes[i];
            if (node.nodeType == 1) {
                var nodeName = node.nodeName;
                switch (nodeName) {
                    case 'A':
                        this.cleanText(node);
                        break;
                    case 'B':
                        var strong = document.createElement('strong');
                        this.replaceTag(strong, node);
                        break;
                    case 'STRONG':
                        this.cleanAttributes(node);
                        this.cleanText(node);
                        break;
                    case 'I':
                        var em = document.createElement('em');
                        this.replaceTag(em, node);
                        break;
                    case 'EM':
                        this.cleanAttributes(node);
                        this.cleanText(node);
                        break;
                    case 'U':
                        this.cleanAttributes(node);
                        this.cleanText(node);
                        break;
                    case 'S':
                        this.cleanAttributes(node);
                        this.cleanText(node);
                        break;
                    case 'SUP':
                        this.cleanAttributes(node);
                        this.cleanText(node);
                        break;
                    case 'SUB':
                        this.cleanAttributes(node);
                        this.cleanText(node);
                        break;
                    case 'BR':
                        this.cleanAttributes(node);
                        debug('clean BR');
                        break;
                    case 'CODE':
                        this.cleanAttributes(node);
                        this.cleanText(node);
                        break;
                    default:
                        if (exclude.indexOf(nodeName) == -1) {
                            this.pushChildrenTop(node);
                        } else {
                            untouched.push(node);
                        }
                }
            } else if (node.nodeType == 3) {
                this.trimText(node);
            }
        }
        return untouched;
    };

    Terrain.prototype.pushChildrenTop = function(node) {
        var parentNode = node.parentNode;
        var childNodes = node.childNodes;
        for (var i = 0; i < childNodes.length; i++) {
            parentNode.insertBefore(childNodes[i], node);
        }
        parentNode.removeChild(node);
    };

    Terrain.prototype.clean = function(node) {
        debug('clean');
        this.cleanRoot(this.dom);
    };

    Terrain.prototype.replaceTag = function(new_node, old_node) {
        var parentNode = old_node.parentNode;
        parentNode.insertBefore(new_node, old_node);

        var n = old_node.childNodes.length;
        var old_nodes = [];
        for (var i = 0; i < n; i++) {
            old_nodes.push(old_node.childNodes[i]);
        }

        for (var i = 0; i < n; i++) {
            new_node.appendChild(old_nodes[i]);
        }

        parentNode.removeChild(old_node);
    };


    Terrain.prototype.cleanRoot = function(root) {
        debug('cleanRoot');
        var n = root.childNodes.length;
        for (var i = n - 1; i >= 0; i--) {
            var node = root.childNodes[i];
            if (node.nodeType == 1) { // TAG
                var nodeName = node.nodeName;
                switch (nodeName) {
                    case 'P':
                        debug('clean P');
                        this.cleanAttributes(node);
                        this.cleanText(node);
                        break;
                    case 'CODE':
                        debug('clean CODE');
                        this.cleanAttributes(node, ['component']);
                        this.cleanText(node);
                        break;
                    case 'BLOCKQUOTE':
                        this.cleanAttributes(node);
                        this.cleanText(node);
                        break;
                    case 'FIGURE':
                        debug('clean FIGURE');
                        this.cleanAttributes(node);
                        break;
                    case 'UL':
                    case 'OL':
                        this.cleanList(node);
                        break;
                    case 'H1':
                        debug('clean H1');
                        this.cleanAttributes(node);
                        this.cleanText(node);
                        var h2 = document.createElement('h2');
                        this.replaceTag(h2, node);
                        break;
                    case 'H2':
                    case 'H3':
                    case 'H4':
                    case 'H5':
                    case 'H6':
                        debug('clean H2');
                        this.cleanAttributes(node);
                        this.cleanText(node);
                        break;
                    case 'TABLE':
                        debug('clean TABLE');
                        this.cleanAttributes(node);
                        break;
                    case 'DIV':
                        debug('clean DIV');
                        //var p = document.createElement('p');
                        //this.replaceTag(p, node);
                        this.cleanRoot(node);
                        this.pushChildrenTop(node);
                        break;
                    default:
                        debug('remove ' + nodeName)
                        root.removeChild(node);
                }
            } else if (node.nodeType == 3) { // TEXT
                if (node.nodeValue != '') {
                    var p = document.createElement('p');
                    root.insertBefore(p, node);
                }
                p.appendChild(node);
            }
        }
    };

    context.Terrain = Terrain;

})(window);
