var profile = {

    releaseDir: "./release",
    releaseName: "v1",
    basePath: './',
    action: 'release',
    mini: true,
    layerOptimize: 'closure',
    stripConsole: 'all',
    selectorEngine: 'acme',
    cssOptimize: 'comments',
    cssImportIgnore: "Calendar.css, ColorPalette.css, Dialog.css, Editor.css, ProgressBar.css, TimePicker.css, Tree.css," +
                     "form/NumberSpinner.css, form/Slider.css, layout/AccordionContainer.css, layout/TabContainer.css",
    
    packages:[
        {
            name: "dojo",
            location: "../dojotoolkit/dojo"
        },
        {
            name: "dijit",
            location: "../dojotoolkit/dijit"
        },
        {
            name: "dojox",
            location: "../dojotoolkit/dojox"
        },
        {
            name: "olesk",
            location: "../olesk"
        }
    ],

    layers: {
        "dojo/dojo": {
            include: ["dojo/dojo"],
            customBase: true,
            boot: true
        },
        "ololo/widget": {
            include: [
                "dojo/request",
                "dojox/atom/io/model",
                "dojo/_base/array",
                "dojo/domReady",
                "dijit/TitlePane",
                "olesk/GoogleBlogger/PostSummary"
            ]
        }
    },

    staticHasFeatures: {
        // The trace & log APIs are used for debugging the loader, so we don’t need them in the build
        'dojo-trace-api': 0,
        'dojo-log-api': 0,
        // This causes normally private loader data to be exposed for debugging, so we don’t need that either
        'dojo-publish-privates': 0,
        // We’re fully async, so get rid of the legacy loader
        'dojo-sync-loader': 0,
        // We aren’t loading tests in production
        'dojo-test-sniff': 0
    }
}
