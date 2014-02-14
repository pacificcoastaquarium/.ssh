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
        }
    ],

    layers: {
        "dojo/dojo": {
            include: ["dojo/dojo"],
            customBase: true,
            boot: true
        },
        "ololo/dojocore": {
            include: [
                "dojo/dom",
                "dojo/dom-attr",
                "dojo/dom-class",
                "dojo/dom-construct",
                "dojo/dom-style",
                "dojo/domReady",
                "dojo/fx",
                "dojo/i18n",
                "dojo/on",
                "dojo/_base/array",
                "dojo/_base/declare",
                "dojo/_base/Deferred",
                "dojo/_base/kernel",
                "dojo/_base/lang",
                "dojo/io/script"
            ]
        },
        "ololo/content_pane": {
            include: [
                "dijit/layout/ContentPane"
            ],
            exclude: ['ololo/dojocore']
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
