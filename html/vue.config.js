const { defineConfig } = require('@vue/cli-service')

module.exports = defineConfig({
  pages: {
    index: {
      entry: 'src/main.ts',
      title: '手機App版本設定介面'
    }
  },
  transpileDependencies: true
})
