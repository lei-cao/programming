<template>
  <div class="about light">
    <aside>
      <header>
        <h1>Tong ZHANG</h1>
      </header>

      <nav>
        <h3>
          Contacts
        </h3>

        <ul>
          <li>
            <div>
              <font-awesome-icon icon="phone"></font-awesome-icon>
              +65 9668 6834
            </div>
          </li>
          <li>
            <div>
              <font-awesome-icon icon="at"></font-awesome-icon>
              taracheung729@gmail.com
            </div>
          </li>
          <li>
            <div>
              <img class="svg-icon" src="/icons/youtube-480.svg"/>
              <a href="https://www.youtube.com/@TaraCooking" target="_blank">@TaraCooking</a>
            </div>
          </li>
          <li>
            <div>
              <img class="svg-icon" src="/icons/linkedin.svg"/>
              <a href="https://www.linkedin.com/in/tara-zhang-0aa04a112" target="_blank">@tara-zhang-0aa04a112</a>
            </div>
          </li>
        </ul>
      </nav>

      <section>
        <h3>SKILLS</h3>

        <ul>
          <li>
            Social Media
          </li>
          <li>
            Digital Marketing
          </li>
          <li>
            Google Ads
          </li>
          <li>
            Google Analytics
          </li>
          <li>
            MS Office
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
             <img class="svg-icon" src="/icons/youtube-480.svg"/> Content Creator  
              <a href="https://www.youtube.com/@TaraCooking" target="_blank">@TaraCooking</a>
            </h4>
            <time>
              since April 2022
            </time>
          </header>

          <summary>
            <ul>
              <li>
                Managing My Personal Youtube Cooking Channel
                <div>
                  - Filming and editing the cooking videos
                </div>
                <div>
                  - Monitoring an engaging with the audiences daily.
                </div>
                <div>
                  - Studying the audiences and choosing the best niches.
                </div>
              </li>
            </ul>
          </summary>
        </article>


        <article>
          <header>
            <h4>
              BUSINESS DEVELOPMENT MANAGER @ LIVE EXPRESS PLATFORM 
            </h4>
            <time>
              Jan 2021 - July 2021
            </time>
          </header>

          <summary>
            Live Express is an online food and grocery delivery platform, which is headquartered in Singapore and operates in SEA countries.
            <ul>
              <li>
                Overachieved the targets at 180% in the first quarter.
              </li>
              <li>
              Closed 60+ new clients which have 120+ branches in total.
              </li>
              <li>
                Monitoring the accounts closely daily and making sure successful onboardings.
              </li>
              <li>
                Upselling the promotions and ads to existing clients and generating more revenues.
              </li>
            </ul>
          </summary>
        </article>

        <article>
          <header>
            <h4>
              Financial Services Consultant @ AIA
            </h4>
            <time>
              Mar 2016 - Oct 2018
            </time>
          </header>

          <summary>
            <ul>
              <li>
                <div>
                  Professional financial executive in managing and supporting sales of insurance products and services.
                </div>
              </li>
              <li>
                <div>
                  Managed more than 200 high net worth Singapore-based and overseas clients.
                </div>
              </li>
              <li>
                <div>
                  Achieved twice MDRT, Top Financial Services Consultant in 2016 and 2017.
                </div>
              </li>
              <li>
                <div>
                  Outstanding success in growing revenues for organisation in managing through effective leads generation and training junior financial consultants, able to exceed sales target up to 150% for the second year running.
                </div>
              </li>
              <li>
                <div>
                  Strong organisational skills with ability to meet deadlines and communicate with key stakeholders and excellent interpersonal and communication skills.
                </div>
              </li>
            </ul>
          </summary>
        </article>

        <article>
          <header>
            <h4>
              Educational Consultant @ EIC Group
            </h4>
            <time>
              Mar 2012 – Jul 2016
            </time>
          </header>

          <summary>
            <ul>
              <li>
                <div>
                  Closed over 1000 deals, won the top monthly sales reward in almost all 4 years.
                </div>
              </li>
              <li>
                <div>
                  Provided personalised advices to the clients for preparing the university applications.
                </div>
              </li>
              <li>
                <div>
                  Provided end-to-end supports to the clients from preparation to the successful applications.
                </div>
              </li>
              <li>
                <div>
                  Provided ongoing training and sharing to other team members.
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
            <h5>
              B.A. in Business Administration and Management @ Geely University of Beijing
            </h5>
            <time>
              Sep 2009 - Jan 2013
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
      header
        h1
          text-align left

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
          margin-bottom 40px
          padding-left 10px
          header
            display flex
            align-items center
            justify-content space-between
            h4
              margin 0
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
        width 380px
        margin-right 20px
        h1
          font-size 25px
        h2
          font-size 16px
          margin-bottom 42px

        ul li
          font-size 12px

      main
        width 900px
        ul li
            font-size 12px
        h4
          font-size 16px
        time
          text-align right
          min-width 115px

    .hide-print
      display none

</style>