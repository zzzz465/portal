module.exports = {
  extends: ['next/core-web-vitals', 'prettier'],
  rules: {
    // A temporary hack related to IDE not resolving correct package.json
    'import/no-extraneous-dependencies': 'off',
    // Since React 17 and typescript 4.1 you can safely disable the rule
    'react/react-in-jsx-scope': 'off',
    '@typescript-eslint/interface-name-prefix': 'off',
    'no-case-declarations': 'error',
    'no-underscore-dangle': 'off',
    'no-restricted-syntax': ['off'],
    '@typescript-eslint/no-namespace': ['off'],
    'func-names': ['off'],
    'react/jsx-props-no-spreading': [
      'warn',
      {
        html: 'enforce',
        custom: 'ignore',
      },
    ],
    'class-methods-use-this': 'off',
    'import/prefer-default-export': 'off',
  },
  parserOptions: {
    ecmaVersion: 2021,
    sourceType: 'module',
    project: './tsconfig.json',
    tsconfigRootDir: __dirname,
    createDefaultProgram: true,
  },
  settings: {
    'import/parsers': {
      '@typescript-eslint/parser': ['.ts', '.tsx'],
    },
  },
}
