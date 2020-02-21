<template>
  <div class="about light">
    <aside>
      <header>
        <h1>Lei CAO</h1>
        <h2>Solutions Engineer</h2>
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
              <font-awesome-icon icon="globe"></font-awesome-icon>
              <a href="https://lei-cao.github.io" target="_blank">lei-cao.github.io</a>
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
            Java, Android, Flutter
          </li>
          <li>
            Python, PHP
          </li>
        </ul>

      </section>

      <section class="links">
        <h3>Projects</h3>

        <ul>
          <li>
            <div>
              <a href="https://remember.gifts" target="_blank">Countdown App</a>
            </div>
          </li>
        </ul>
      </section>


      <section class="links hide-print">
        <h3>Links</h3>

        <ul>
          <li>
            <font-awesome-icon icon="file-download"></font-awesome-icon>
            <a href="/files/lei-cao-resume.pdf">Download Resume</a>
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
              Senior Solutions Engineer @ Remerge
            </h4>
            <time>
              since May 2019
            </time>
          </header>

          <summary>
            <ul>
              <li>
                Being the technical expert and reference, working closely with sales, account managers and clients,
                managing the tech integrations for APAC clients such as Shopee, Oyorooms, Sendo, Traveloka, Grab, Clubfactory, Firecraft, etc.
              </li>
              <li>
                <div>
                  Built Template Manager which is used by designers and solutions engineers to create ads banner templates.
                  It streamlined the dynamic template creation process and improved efficiency significantly.
                </div>
                <div>
                  Tech stack: Go, ReactJS, Firebase
                </div>
              </li>
              <li>
                Maintain and improve the integrations with partners such as AppsFlyer, Adjust, Branch.
              </li>
              <li>
                Working part-time as a product manager leading the operation efficiency project.
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
                  Led the App integration efficiency project by creating automated processes and tools.
                  Dug into the details of the processes to find the bottlenecks, built the products that solve
                  the problems in scale.
                  As the main developer, leading the team of 3 engineers for coding and delivery the tools.
                </div>
                <div>
                  Tech stack: Go, Vue, WebSocket, MySQL, Protobuf, Metabase, Python, Mitmproxy.
                </div>
              </li>
              <li>
                Worked with Product Team and R&D closely for product improvements.
              </li>
              <li>
                Worked with Mobile Measurement Partners such as AppsFlyer, Adjust, Tune and Branch, being the internal technical
                expert for these partners.
              </li>
              <li>
                Been the technical expert and reference for App products in the APAC region.
                Built documentation and training materials, conducted training to TS, commercial team and clients.
              </li>
              <li>
                Managed App integrations of top clients in SEA such as Lazada, Shopee, Zalora, Hotels, Agoda, Traveloka.
                Worked with sales and account managers through whole customer journeys.
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
                  Led the Mobile API project. Redesigned and built the Mobile API in Go. Improved the API performance
                  and stability significantly.
                  The API was used by iOS and Android app, React native and mobile web app.
                </div>
                <div>
                  Tech stack: Go, RabbitMQ, Couchbase, ElasticSearch, InfluxDB, MySQL, Thrift, Docker.
                </div>
              </li>
              <li>
                Contributed to open source project Beego, internationalized the official documentation in English.
                Advocated for the project by creating articles and video tutorials.
              </li>
              <li>
                As the sole developer, designed and developed a Zendesk App to connect the customers with profile and
                order information for the customer service team.
              </li>
              <li>
                <div>
                  Maintained Zalora desktop and mobile e-commerce websites.
                </div>
                <div>
                  Tech stack: PHP, JQuery, AngularJS, MySQL, Solr, Selenium
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
                  As one of the 2 developers, with another 3 co-founders, we created a knowledge collecting and sharing
                  community. Designed the product and built the whole web app from scratch.
                </div>
                <div>
                  Tech stack: Python, Flask, JQuery, Solr, Redis, MySQL.
                </div>
              </li>

              <li>
                <div>
                  Built a consumer product recommendation and review platform. Worked as the sole Android Developer and built the Android app from scratch.
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
        font-size 42px
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