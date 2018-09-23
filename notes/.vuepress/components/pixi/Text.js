import Sprite from './Sprite.js'

export default {
  mixins: [Sprite],
  props: {
    text: String
  },
  computed: {
    instance: () => new PIXI.Text()
  },
  watch: {
    'instance': {
      handler (newInstance, oldInstance) {
        if (this.text) newInstance.text = this.text
      },
      immediate: true
    },
    'text': function (text) { this.instance.text = text }
  }
}
