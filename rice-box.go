package main

import (
	"github.com/GeertJohan/go.rice/embedded"
	"time"
)

func init() {

	// define files
	file2 := &embedded.EmbeddedFile{
		Filename:    "app.js",
		FileModTime: time.Unix(1498303041, 0),
		Content:     string("document.getElementById(\"text\").innerHTML = \"Hello World<br /><p>Created by Nic Raboy</p>\";"),
	}
	file3 := &embedded.EmbeddedFile{
		Filename:    "custom.css",
		FileModTime: time.Unix(1498303034, 0),
		Content:     string(".container {\r\n    width: 400px;\r\n    height: 100px;\r\n    background-color: #009ACD;\r\n    margin: auto;\r\n}\r\n \r\n.text {\r\n    position: relative;\r\n    top: 50%;\r\n    transform: translateY(-50%);\r\n    text-align: center;\r\n    color: #FFFFFF;\r\n}"),
	}
	file4 := &embedded.EmbeddedFile{
		Filename:    "index.html",
		FileModTime: time.Unix(1498303048, 0),
		Content:     string("<html>\r\n    <head>\r\n        <title>The Polyglot Developer</title>\r\n        <link rel=\"stylesheet\" href=\"custom.css\" />\r\n    </head>\r\n    <body>\r\n        <div class=\"container\">\r\n            <p id=\"text\" class=\"text\"></p>\r\n        </div>\r\n        <script src=\"app.js\"></script>\r\n    </body>\r\n</html>"),
	}

	// define dirs
	dir1 := &embedded.EmbeddedDir{
		Filename:   "",
		DirModTime: time.Unix(1498302910, 0),
		ChildFiles: []*embedded.EmbeddedFile{
			file2, // "app.js"
			file3, // "custom.css"
			file4, // "index.html"

		},
	}

	// link ChildDirs
	dir1.ChildDirs = []*embedded.EmbeddedDir{}

	// register embeddedBox
	embedded.RegisterEmbeddedBox(`website`, &embedded.EmbeddedBox{
		Name: `website`,
		Time: time.Unix(1498302910, 0),
		Dirs: map[string]*embedded.EmbeddedDir{
			"": dir1,
		},
		Files: map[string]*embedded.EmbeddedFile{
			"app.js":     file2,
			"custom.css": file3,
			"index.html": file4,
		},
	})
}
