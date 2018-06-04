
/*
[[INCLUDE component=WikiEditBar]]
[[INCLUDE component=Terrain]]
[[INCLUDE component=WikiEditButtons]]
*/

(function(context) {

    var debug = false;

    var WikiTerm = function(dom) {
        this.dom = dom;
        this.hash = dom.dataset.hash;

        this.title = dom.querySelector('.title');
        this.content = dom.querySelector('.content');
        this.footer = dom.querySelector('.footer');

        this.terrain = null;

        var that = this;
    };

    WikiTerm.prototype.edit = function(callback) {
        wikaan.lockTerm(this.hash, function(response){
            if (response.acquired) {
                that.enable_editor(response.content);
            }
            callback(response.acquired);
        });

        var that = this;
    };

    WikiTerm.prototype.publish = function(callback) {
        var content = this.terrain.dom.innerHTML;

        var final_document = document.createElement('div');
        final_document.innerHTML = content;

        var terrain = new Terrain(final_document);
        terrain.clean();
        terrain.removeEmptyParagraphs();
        terrain.highlightCodes();
        terrain.generateIndex();

        var html = final_document.innerHTML;

        // TODO: publish:
        var that = this;
        console.log('TODO: publish', this.hash, content, html);
        that.content.innerHTML = html;
        that.disable_editor();
        callback(true);

    };

    WikiTerm.prototype.discard = function(callback) {
        wikaan.discard(term.hash, function(result){
            term.disable_editor();
            callback(true);
        });
    };

    WikiTerm.prototype.disable_editor = function() {
        this.am_i_editing = this.dom.dataset.am_i_editing = false;
        this.dom.removeChild(this.terrain.dom);
        this.terrain = null;
        this.dom.removeChild(this.bar.dom);
        this.bar = null;
    };

    WikiTerm.prototype.enable_editor = function(content) {
        this.am_i_editing = this.dom.dataset.am_i_editing = true;

        var editor = document.createElement('div');
        editor.innerHTML = content;
        editor.classList.add('editor');
        this.dom.appendChild(editor);

        var terrain = this.terrain = new Terrain(editor);
        this.terrain.enableEditor();

        var upload_file = function(file, link) {
            link.setAttribute('uploading', true);

            var formData = new FormData();
            formData.append("file", file);
            formData.append("wiki_hash", that.hash);

            var xhr = new XMLHttpRequest();
            xhr.addEventListener('load', function(e){
                link.removeAttribute('uploading');
                var json = JSON.parse(xhr.response);
                link.href = 'files/' + json.hash;
            }, false);
            xhr.upload.onprogress = function(e) {
                // var p = 100 * e.loaded / e.total;
            };
            xhr.open('post', '[[AJAX name=upload_file]]', true);
            xhr.send(formData);
        };

        var upload_image = function(file, img) {
            var formData = new FormData();
            formData.append("file", file);
            formData.append("wiki_hash", that.hash);

            var xhr = new XMLHttpRequest();

            xhr.addEventListener("load", function(e) {
                var json = JSON.parse(xhr.response);
                img.src = 'images/' + json.hash + '/w:400px;';
            }, false);
            xhr.upload.onprogress = function(e) {
                // var p = 100 * e.loaded / e.total;
            };
            xhr.open('post', '[[AJAX name=upload_image]]', true);
            xhr.send(formData);
        };

        var input_file = document.createElement('input');
        input_file.setAttribute('type', 'file');
        input_file.setAttribute('multiple', true);
        input_file.addEventListener('change', function(e) {
            var files = input_file.files;
            for (var i=0; i<files.length; i++) {
                if (i>0) {
                    terrain.insertElement(document.createTextNode(', '));
                }
                var file = files[i];
                var link = terrain.insertFileLink('#', file.name, file.type);
                upload_file(file, link);
            }
        }, true);

        var input_image = document.createElement('input');
        input_image.setAttribute('type', 'file');
        input_image.setAttribute('multiple', true);
        input_image.addEventListener('change', function(e) {
            var files = input_image.files;
            for (var i=0; i<files.length; i++) {
                var img = terrain.insertImage('data:image/png;base64,iVBORw0KGgoAAAANSUhEUgAAAQAAAAEACAYAAABccqhmAAALUUlEQVR42u3cf2zUdx3H8df7vr3rUQYrHR10LtNtyqJIGCvi4pz+g2OLTqeRRQwyDJG767Za2fyNpsmSyWKyNMy238MZFWHdkjk2lszFZP+NGYwdSySOsaFT4yqI4AaFXr987+MffG/5etz1x4EHLc9HQnLX3n2/1++3n+d9Pt+2SAAA4OJjZ/PkfD4/U1KSwwicN0EmkxmuWwDKBv0qSQs4B8B5s1/SQK0xsBoGfvmg7+IcAOdNT4UYTDgENonBvyY28LvGeBEA6qPaONwvaetEImCTHPxd45QHQH2MNRPvmWgErIbB33Mu1h4AajfGtbiuyUTAahj8k15nAKhLDFZVGa9VI2A1DP6tDHzggg3BpMatMfiBizcCiQrbSDL4gakpGqNbozHbE7susEAVfmkvUaEeq2IfYvAD0yMCkrQqGuNVZwDl7/6SNMDgB6ZkBOI/mq84C0iMsY3Su3/A4QSmpKDCLKDyEqDC9J93f2B6zQLOWAYkxpn+A5g+zlgGJJj+AxfvMmCsawBM/4HpuQyYUAAATHMEACAAAAgAAAIAgAAAIAAACAAAAgCAAAAgAAAIAAACAIAAACAAAAgAAAIAgAAAIAAACAAAAgCAAAAgAAAIAAACAIAAACAAAAgAAAIAgAAAIAAACAAAAgCAAAAgAAAIAAACAIAAAASAQwAQAAAEAAABAEAAABAAAAQAAAEAQAAAEAAABAAAAQBAAAAQAAAEAAABAEAAABAAAAQAAAEAQAAAEAAABAAAAQBAAAAQAAAEAAABAEAAABAAAAQAAAEAQAAAEAAABAAAAQBAAAACAIAAACAAAAgAAAIAgAAAIAAACAAAAgCAAAAgAAAIAAACAIAAACAAAAgAAAIAgAAAIAAACAAAAgCAAACojwYOQX3k83lXup3JZIwjAmYAAAgAAAIAgAAAqBcuAl4gfN//kqSlkhaZ2dWS5jnnZphZUdKQpBeCIHjonnvueX3z5s3XpFKptWa2wjl3jaTZkg5L+r2kLdls9jfV9pPP53/onLtZ0lVm1uqcu8TMjjrnBovF4q8OHTr0eHd3t6v2/N7e3pmJRGKtmd1mZtc75y6L3khOSnpL0gEzeyUMwz90dHQ8U2kbDz/88JympqaMmX1G0oecc5dIOiLpFUlPjo6Obu3s7Bzlu4IZwEXDzAbM7D4zu0XSByTNNrOkpEZJ75O0LplMvuz7/vZUKrXfzH4gaZmZzTWzlJldYWZ3mNlz+Xx+0xi7usvMlpvZAklzon1cbma3eZ732Pz585/o7u62KpH6rOd5byYSiZ+Y2aclvcfM0tH+LzWzD0aDeqPneU9X2kZfX9/nmpqa/mxmP5J0U+k1mNk8M1thZj9NpVK7+/r6ruS7ggBclJxzS4vF4uVDQ0MNw8PDLcViMeecCyRdYmZfNrN3isXid8IwvG5oaCgdhuEc59zXJBWiTXy7t7f3pirb/n6xWLy5UCi0Dg4OJoeHh1ui556MQrRy/vz56yoM3FWSnjazuZJOOuceCIJg4dDQUHpwcDBZKBRawzC8cayvq6+v73bP854ys2ZJR51z6wqFQuvQ0FA6CIIbnHM7o9dwved5O9avX88MlSXAxSebzQ7G7h6V5Pf3919rZvdHH3s2l8s9FHtMQdKj+Xz+vZI2SlJDQ8NXJe2qsO3Hyz50NHpuq6QHowH4FUmPlh7Q399/hZk9ambmnAucc7fkcrkXy7ZzWNLhfD5f8Wvq6em51PO8n0tKOOfeDoLghnvvvffN2EP2dHd339HW1vacpFslLV2yZMmdkh7jO4IZALMC556M3W6v9JgwDHfEHrNkMtsvFApPxO4uKluefN3MmqLbj1QY/ONKp9MZSZdFdx8oG/ySpO7ublcsFh+I7feLnHkCAEknTpx4I3a3rdJjRkZG/hIbPPMms/0jR478PRaP5rJP3xqLzPYav4TbSzdGR0d3VHvQyZMnX47dbefMswSApOPHj78za9as0uBurvSYffv2HWtvby8N4tZKj/F9f4Wk1ZKWSbrSzFKSDjrn9sbiUX4R8OrSjeHh4X01fgnXlW40NjYeqLZUKDOXM88MAKenx8F4523Lli2nYoM4Ff/c5s2bU77vP2Vmz5vZajNbEE3rG3T6av6KMXafjAUgrPFLaK7hOTM488wAcA6kUqmNZvb5aHZwxDn3rSAInt+7d+/BxYsXN5rZdZ7nDVZ5+iFJV0nS3Llzr5L0eg0v4ZikligiLRs2bDjKWSEAqJ81sTX+7blc7qXY505JernatNzMdpUC4HneHZJ+XMP+90u6UZLS6fQNkl7glLAEQP28e+Hw4MGDeybzxDAMf/buN0wi8V3f999fw/6fL93wPG8Np4MZAOrrn6V38Xnz5i2R9NJEn9jR0fGC7/vbzGy1pDmSXurv7984MjKy89VXXz28aNGi1mQy+TEzy1XbxsjISH86nf6mmc2UtNr3/V9ns9mdnBYCgDowsyclbYhub/V9f8Px48d37d69++jChQsbW1pa2sZ6/ujo6LpUKhWa2V1m1mpm+aampnzppw7j6erqOpTP5zOStkUzz6d8398ShuH2IAj27tq169iyZcvSM2fOvNLzvI8Ui8WWXC73CGeOJQDOgUKh0C3pd1EArjWzZ2bNmnV4+fLlYVtb24nGxsYDYz2/s7NzNJvNrj116tTHnXO/dM69odO/PlyQ9KZzbptz7lOx6wxn/EFRJpPZ7pz7gqR/m5lnZrmGhoYXZ8yY8Z/ly5eHs2fPHvY87zVJ28zsG5w1ZgA4Rzo7O4+tX7/+E+3t7audc3dKWiypVVLCzE46547o9F8d/lXSa9W2c/fdd+9ShV8xlqRNmzbNnjNnTunu25Uek81md/T29v62oaHhLufcCknXS5prZinn3HFJf5O0x8ye46wRgGljvP8HcCL/T+DZPib6PYFfRP/Ouebm5qWxuwfGiMiwpL7oH1gCYKpbuXJlwsy+F/vQsxwVZgCYJnzfv8/Mmp1zLxYKhT+FYfivffv2nVq8ePFlnud9VNL9km6O1v9vjYyM9HDUCACmCTNbK+nDZqZ0Oi1JqvQTAOfcH4MgWNnV1fU2R40AYJoIw/DBRCLxSZ3+U+F5ZtYqqSm6cPcPSXucczv27NmzM/43CSAAmAY6OjoGJA1wJKYXLgICBAAAAQBAAAAQAAAEAAABAEAAABAAAAQAAAEAQAAAEAAABAAAAQBAAAAQAAAEAAABAEAAABAAAAQAAAEAQAAAEAAABAAAAQBAAAAQAAAEAAABAEAAABAAAAQAAAEAQAAAEACAAHAIAAIAgAAAIAAACAAAAgCAAAAgAAAIAAACAIAAACAAAAgAAAIAgAAAIAAACAAAAgCAAAAgAAAIAAACAIAAACAAAAgAAAIAgAAAIAAACAAAAgCAAAAgAAAIAAACAIAAACAAAAgAQAAAEAAABAAAAQBAAAAQAAAEAAABAEAAABAAAAQAAAEAQAAAEAAABAAAAQBAAAAQAAAEAAABAEAAABAAAOc/AKvy+fxMDhEwdUVjeNVkA9AlaYGkJIcQmNKS0VjuGi8AgaT9kno4ZsC01BON8eCMAGQymWFJAywDgGk9/R+Ixvq41wBYBgDTePpfKQCVlgHMAoCp/+5/xvT/jABUWAaUZgFriAAwpQb/mgrv/v8z/a+2BCifBRABYGoP/orv/pJkNWxka3lFAFzwg7/iuLUaNzYgKSAEwAUz8JPRmn9Sb9pWY1EUC4GIAXDeBr1iA1+TnbHbWVxQiP+kIB4DAP9/8UGvCmNzQst1m8ieyiJQvrPyGACoj2rjcMLX6myie6qwzqj2IgDUT6WZ+ISX5DbZvY2x9gBQf2d1Lc7OZs9lMQBQf1yABwAAk/RflmGkbyIISp0AAAAASUVORK5CYII=');
                var file = files[i];
                upload_image(file, img);
            }
        }, true);

        var bar = this.bar = new WikiEditBar();
        this.dom.insertBefore(bar.dom, this.title);

        bar.addGroup('title')
            .addButtons('title')
            .addSlideButtons('title-1', function(event) {
                terrain.formatBlock('H2');
            })
            .addSlideButtons('title-2', function(event) {
                terrain.formatBlock('H3');
            })
            .addSlideButtons('title-3', function(event) {
                terrain.formatBlock('H4');
            })
            .addSlideButtons('title-4', function(event) {
                terrain.formatBlock('H5');
            })
            .addSlideButtons('title-5', function(event) {
                terrain.formatBlock('H6');
            });

        bar.addGroup('text')
            .addButtons('bold', function(event) {
                terrain.formatBold();
            })
            .addButtons('italic', function(event) {
                terrain.formatItalic();
            })
            .addButtons('underline', function(event) {
                terrain.formatUnderline();
            })
            .addButtons('strike', function(event) {
                terrain.formatStrike();
            })
            .addButtons('remove-format', function(event) {
                // TODO: replace this with a terrain function
                document.execCommand('removeFormat', false, null);
            });

        bar.addGroup('lists')
            .addButtons('ordered-list', function(event) {
                terrain.insertOl();
            })
            .addButtons('unordered-list', function(event) {
                terrain.insertUl();
            })
        ;

        bar.addGroup('script')
            .addButtons('superscript', function(event) {
                // TODO: replace this with a terrain function
                document.execCommand('superscript', false, null);
            })
            .addButtons('subscript', function(event) {
                // TODO: replace this with a terrain function
                document.execCommand('subscript', false, null);
            })
        ;

        bar.addGroup('blocks')
            .addButtons('quotes', function(event) {
                terrain.insertQuotes();
            })
            .addButtons('code', function(event) {
                terrain.insertCode();
            })
            /*
            .addButtons('image', function(event) {
                input_image.click();
            })
            //*/
        ;

        /*
        bar.addGroup('links')
            .addButtons('file', function(event) {
                input_file.click();
            })
            .addButtons('href', function(event) {
                var href = prompt('Link URL, for example: http://www.google.com/');
                if (null === href) {
                    return;
                }

                terrain.insertExternalLink(href);
            })
        ;
        //*/

        var big_buttons = bar.addGroup('big_buttons');
        big_buttons.addButton('publish', function(event) {
            context.term.publish(function(a,b,c) {

				var xhr = new XMLHttpRequest();
				xhr.open('POST', '/v0/articles', true);
				xhr.onload = function() {
					window.location.href = '/';
				};

				var payload = {
					"title": document.getElementById('article-title').textContent,
					"content": editor.innerHTML,
				};

				xhr.send(JSON.stringify(payload));
			});
        }).setText('Publish');

        editor.addEventListener('blur', function(e) {
            console.log('TODO: save draft', that.hash, editor.innerHTML);
        }, true);

        var that = this;
    };

    window.addEventListener('load', function() {
        component = document.querySelector('[data-component="WikiTerm"]');
        if (null !== component) {
            console.log(component);
            context.term = new WikiTerm(component);
            context.term.enable_editor('');
        }
    }, true);

})(window);
