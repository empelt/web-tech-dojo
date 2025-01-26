import { FlatCompat } from '@eslint/eslintrc'
import pluginJs from '@eslint/js'
import prettierConfig from 'eslint-config-prettier'
import * as importPlugin from 'eslint-plugin-import'
import perfectionistPlugin from 'eslint-plugin-perfectionist'
import pluginReact from 'eslint-plugin-react'
import unusedImportsPlugin from 'eslint-plugin-unused-imports'
import globals from 'globals'
import tseslint from 'typescript-eslint'

const flatCompat = new FlatCompat()

/** @type {import('eslint').Linter.Config[]} */
export default [
  { settings: { react: { version: 'detect' } } },
  { ignores: ['dist'] },
  { files: ['**/*.{js,mjs,cjs,ts,jsx,tsx}'] },
  { languageOptions: { globals: globals.browser } },
  pluginJs.configs.recommended,
  ...tseslint.configs.strict,
  pluginReact.configs.flat.recommended,
  ...flatCompat.extends('plugin:react-hooks/recommended'),
  {
    rules: {
      'react/destructuring-assignment': 'error',
      'react/function-component-definition': [
        'error',
        {
          namedComponents: 'arrow-function',
          unnamedComponents: 'arrow-function',
        },
      ],
      'react/hook-use-state': 'error',
      'react/jsx-fragments': 'error',
      'react/jsx-curly-brace-presence': 'error',
      'react/jsx-no-useless-fragment': 'error',
      'react/jsx-sort-props': 'error',
      'react/self-closing-comp': 'error',
      'react/jsx-pascal-case': 'error',
      'react/no-danger': 'error',
      'react/prop-types': 'off',
    },
  },
  {
    plugins: { import: importPlugin },
    rules: {
      'import/order': [
        'warn',
        {
          groups: [
            'builtin',
            'external',
            'internal',
            ['parent', 'sibling'],
            'object',
            'type',
            'index',
          ],
          'newlines-between': 'always',
          pathGroupsExcludedImportTypes: ['builtin'],
          alphabetize: { order: 'asc', caseInsensitive: true },
          pathGroups: [
            {
              pattern: 'react',
              group: 'external',
              position: 'before',
            },
          ],
        },
      ],
    },
  },
  {
    plugins: { 'unused-imports': unusedImportsPlugin },
    rules: {
      '@typescript-eslint/no-unused-vars': 'off',
      'unused-imports/no-unused-imports': 'error',
      'unused-imports/no-unused-vars': [
        'error',
        {
          vars: 'all',
          varsIgnorePattern: '^_',
          args: 'after-used',
          argsIgnorePattern: '^_',
        },
      ],
    },
  },
  {
    plugins: { perfectionist: perfectionistPlugin },
    rules: {
      'perfectionist/sort-interfaces': 'warn',
      'perfectionist/sort-object-types': 'warn',
    },
  },
  prettierConfig,
  {
    rules: {
      'react/jsx-uses-react': 'off',
      'react/react-in-jsx-scope': 'off',
    },
  },
]
