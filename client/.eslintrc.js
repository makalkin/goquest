module.exports = {
    "extends": ["standard", "plugin:react/recommended"],
    "plugins": [
        "standard",
        "promise",
        "react"
    ],
    "env": {
        "es6": true,
        "browser": true,
        "node": true
    },
    "parser": "babel-eslint",
    "parserOptions": {
        "ecmaVersion": 6,
        "sourceType": "module",
        "ecmaFeatures": {
            "jsx": true,
            "experimentalObjectRestSpread": true,
            "spread": true
        }
    },
    "settings": {
        "react": {
            "pragma": "React",  // Pragma to use, default to "React"
            "version": "15.0" // React version, default to the latest React stable release
        }
    },
    "rules": {
        "semi": ["error", "always"],
        "quotes": ["error", "double"],
        "indent": ["error", 4],
        "no-unused-vars": 0,
        "react/jsx-uses-vars": "error",
        "eol-last": 0,
        "no-class-assign": 0,
        "react/prop-types": 0,
        "react/react-in-jsx-scope": 0,
        "react/no-danger": 0
    }
};