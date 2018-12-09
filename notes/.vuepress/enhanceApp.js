import {library} from '@fortawesome/fontawesome-svg-core'
import {faPhone, faAt, faGlobe, faCircle, faFileDownload} from '@fortawesome/free-solid-svg-icons'
import {faGithub} from '@fortawesome/free-brands-svg-icons'
import {FontAwesomeIcon} from '@fortawesome/vue-fontawesome'

export default ({
                  Vue, // the version of Vue being used in the VuePress app
                  options, // the options for the root Vue instance
                  router, // the router instance for the app
                  siteData // site metadata
                }) => {
  library.add(faPhone, faAt, faGlobe, faCircle, faFileDownload, faGithub)
  Vue.component('font-awesome-icon', FontAwesomeIcon)
}