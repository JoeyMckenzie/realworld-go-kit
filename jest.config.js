module.exports = {
  moduleNameMapper: {
    '@core/(.*)': '<rootDir>/conduit-web/app/core/$1',
  },
  preset: 'jest-preset-angular',
  setupFilesAfterEnv: ['<rootDir>/setup-jest.ts'],
};
