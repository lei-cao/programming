<template>
  <div class="about light">
    <aside>
      <header>
        <h1>Lei CAO</h1>
        <h2>Solutions Engineer Lead</h2>      
      </header>

      <nav>
        <h3>
          Contacts
        </h3>

        <ul>
          <li>
            <div>
              <font-awesome-icon icon="phone"></font-awesome-icon>
              +65 81321483
            </div>
          </li>
          <li>
            <div>
              <font-awesome-icon icon="at"></font-awesome-icon>
              lei.cao.life@gmail.com
            </div>
          </li>
          <li>
            <div>
              <font-awesome-icon :icon="['fab', 'github']"/>
              <a href="https://github.com/lei-cao" target="_blank">github.com/lei-cao</a>
            </div>
          </li>
          <li>
            <div>
              <img class="svg-icon" src="/icons/linkedin.svg"/>
              <a href="https://www.linkedin.com/in/leicaogo/" target="_blank">@leicaogo</a>
            </div>
          </li>
        </ul>
      </nav>

      <section>
        <h3>SKILLS</h3>

        <ul>
          <li>
            Javascript, Vue, React
          </li>
          <li>
            Firebase
          </li>
          <li>
            Go, Docker
          </li>
          <li>
            Android, Flutter
          </li>
          <li>
            Python, PHP, SQL
          </li>
        </ul>

      </section>

      <section class="links">
        <h3>Projects</h3>

        <ul>
          <li>
            <div>
              <a href="https://pro.thinking.poker" target="_blank">Poker Tools</a>
            </div>
          </li>
          <li>
            <div>
              <a href="https://www.youtube.com/@realThinkingPoker" target="_blank">Youtube Channel</a>
            </div>
          </li>
        </ul>
      </section>


      <section class="links hide-print">
        <h3>Links</h3>

        <ul>
          <li>
            <font-awesome-icon icon="file-download"></font-awesome-icon>
            <a href="/files/lei-cao-resume-2023.pdf">Download Resume</a>
          </li>
        </ul>
      </section>

    </aside>

    <main>
      <section>
        <h2>
          Experiences
        </h2>

        <article>
          <header>
            <h4>
              Senior Parter Product Manager @ AppLovin
            </h4>
            <time>
              since June 2022
            </time>
          </header>

          <summary>
            <ul>
              <li>
                Collaborate with sales, product, engineering, and partners to manage the entire product life cycle
              </li>
              <li>
                Work with top mobile device OEMs and carriers to integrate AppLovin products into their devices
              </li>
              <li>
                Utilize BigQuery, Looker, and App Scripts to monitor product performance and for custom reporting
              </li>
            </ul>
          </summary>
        </article>

        <article>
          <header>
            <h4>
              Solutions Engineer Lead @ Remerge
            </h4>
            <time>
              May 2019 - June 2022
            </time>
          </header>

          <summary>
            <ul>
              <li>
                Successfully led a team of three professionals, driving success for complex tech integrations for APAC-based clients including Shopee, Grab, GSShop, Firecraft, AliExpress, Mytona
              </li>
              <li>
                <div>
                  Led the development of Template Manager, streamlining the dynamic template creation process and enabling efficient creation of high-quality ad banners    
                </div>
                <div>
                  Tech stack: Go, ReactJS, Firebase
                </div>
              </li>
            </ul>
          </summary>
        </article>

        <article>
          <header>
            <h4>
              Senior Solutions Engineer @ Criteo
            </h4>
            <time>
              March 2016 - April 2019
            </time>
          </header>

          <summary>
            <ul>
              <li>
                <div>
                  Developed a dashboard for improving app integration efficiency, utilizing ElasticSearch logs with a user-friendly UI, fast speed, and monitoring/alerting capabilities. Widely adopted by solutions engineers as a daily-use tool
                </div>
                <div>
                  Tech stack: Go, Vue, WebSocket, ElasticSearch, Kibana, MySQL, Metabase, Mitmproxy
                </div>
              </li>
              <li>
                Collaborated closely with product team and R&D to drive product improvements
              </li>
              <li>
                Served as internal technical expert for Mobile Measurement Partners including Adjust, AppsFlyer, Kochava, Singular
              </li>
              <li>
                Built documentation and conducted training sessions for clients and team members
              </li>
            </ul>
          </summary>
        </article>

        <article>
          <header>
            <h4>
              Senior Software Engineer @ Zalora
            </h4>
            <time>
              February 2013 - February 2016
            </time>
          </header>

          <summary>
            <ul>
              <li>
                <div>
                  Rebuilt Mobile API in Go, improving performance and stability
                </div>
              </li>
              <li>
                Contributed to open-source project Beego by internationalizing official documentation
              </li>
              <li>
                Developed Zendesk app to connect customers with customer service team with detailed order information
              </li>
              <li>
                <div>
                  Responsible for maintaining Zalora e-commerce websites, both desktop and mobile versions
                </div>
                <div>
                  Tech stack: Go, PHP, JQuery, MySQL, RabbitMQ, Docker, Solr, Selenium
                </div>
              </li>
            </ul>
          </summary>

        </article>

        <article>
          <header>
            <h4>
              Software Engineer, co-founder @ Baygle
            </h4>
            <time>
              September 2009 - January 2013
            </time>
          </header>

          <summary>
            <ul>
              <li>
                <div>
                  Co-created a knowledge collecting and sharing community with 3 co-founders
                </div>
                <div>
                  Designed and built the whole web app from scratch using Python, Flask, JQuery, Solr, Redis, MySQL
                </div>
              </li>
            </ul>
          </summary>
        </article>

        
      </section>

      <section>
        <h2>
          Education
        </h2>

        <article>
          <header>
            <h4>
              B.E. in Automation @ Sun Yat-Sen University, Guangzhou
            </h4>
            <time>
              September 2005 - January 2009
            </time>
          </header>
        </article>

      </section>
    </main>
  </div>
</template>

<script>
  export default {
    name: 'AboutLayout'
  }
</script>

<style lang="stylus" scoped>
  base03 = #002b36
  base02 = #073642
  base01 = #586e75
  base00 = #657b83
  base0 = #839496
  base1 = #93a1a1
  base2 = #eee8d5
  base3 = #fdf6e3
  yellow = #b58900
  orange = #cb4b16
  red = #dc322f
  magenta = #d33682
  violet = #6c71c4
  blue = #268bd2
  cyan = #2aa198
  green = #859900

  rebase(rebase03, rebase02, rebase01, rebase00, rebase0, rebase1, rebase2, rebase3)
    background-color rebase03
    color rebase0
    *
      color rebase0

    h1, h2, h3, h4, h5, h6
      color rebase1
      border-color rebase0

    a, a:active, a:visited
      color rebase1

  accentize(accent)
    a, a:active, a:visited, code.url
      color accent

    h1, h2, h3, h4, h5, h6
      color accent

  html, .light
    rebase(base3, base2, base1, base0, base00, base01, base02, base03)

  .dark
    rebase(base03, base02, base01, base00, base0, base1, base2, base3)

  html *
    color-profile: sRGB
    rendering-intent: auto

  p, ul, ol
    line-height 1.7

  .about
    font-family: -apple-system, BlinkMacSystemFont, "Segoe UI", Roboto, Oxygen, Ubuntu, Cantarell, "Fira Sans", "Droid Sans", "Helvetica Neue", sans-serif;
    -webkit-font-smoothing: antialiased;
    -moz-osx-font-smoothing: grayscale;
    font-size: 16px;

    display flex
    justify-content center

    .svg-icon
      width 16px
      height 16px

    a
      color base00
    a:visited
      color yellow
    a:hover
      color red

    h2, h3, h4, h5
      text-transform uppercase

    aside

      nav
        margin-bottom 35px

      h1, h2
        text-align center

      h1
        font-size 38px
        color base01

      h2
        font-size 22px
        margin-bottom 42px

      h3
        color base0
        background-color base2
        padding 8px

      ul
        list-style-type none
        list-style-position outside
        padding-inline-start 10px

        li
          line-height 30px

      section:not(.links) ul

        li
          line-height 30px

        li:before
          content "·"
          font-size 50px
          vertical-align text-bottom
          line-height 25px

      section.links ul
        li

    main

      section
        h2
          color base0
          background-color base2
          padding 8px

        article
          padding-left 10px
          header
            display flex
            align-items center
            justify-content space-between

          time
            color base1

  /* Portrait tablet to landscape and desktop */
  @media screen and (min-width: 768px)
    .about
      aside
        width 300px
        margin-right 20px

      main
        width 850px

  /* Landscape phone to portrait tablet */
  @media screen and (max-width: 767px)
    .about
      flex-direction column

  @media print

    body
      -webkit-print-color-adjust exact

    .about
      font-size: 11px

      aside
        width 250px
        margin-right 20px
        h2
          font-size 14px
          margin-bottom 42px

        ul li
          font-size 12px

      main
        width 900px

    .hide-print
      display none

</style>