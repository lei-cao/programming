<template>
      <section>
        <article>

            <div>
I'm excited to apply for the Technical Solutions, Team Lead (APAC) role at StackAdapt. With a rich career spanning across various technical roles in the software industry, my unique blend of product management, solutions engineering, leadership, and software development experience position me as a strong candidate for this role.
</div>
<div>

In my current role as Senior Partner Product Manager at AppLovin, I collaborate with multiple stakeholders to manage the product lifecycle, working closely with leading mobile device OEMs including Oppo, Vivo, Xiaomi, Samsung and Realme and carriers to integrate our products. A key part of my role involves leveraging BigQuery, Looker, and App Scripts for product performance monitoring and custom reporting, aligning with the qualifications mentioned in your job description.
</div>
<div>

Prior to this, I held a leadership role at Remerge as a Solutions Engineer Lead. Here, I successfully led a team focused on complex tech integrations for APAC-based clients including prominent names like Shopee, Grab, and AliExpress. I developed the Template Manager using Go, React.js and firebase, streamlining the dynamic template creation process, thereby enhancing the efficiency of ad banner creation. This experience honed my technical leadership and problem-solving skills.
</div>
<div>

During my Senior Solutions Engineer role at Criteo, I had the opportunity to work with multiple technologies like Go, Vue, WebSocket, ElasticSearch, and more. I developed a user-friendly dashboard that improved app integration efficiency and served as an internal technical expert for Mobile Measurement Partners.
</div>
<div>

With strong technical skills, leadership, and customer support experiences, I'm confident of my fit for this role. I look forward to the opportunity to discuss my suitability further.
</div>
<div>

Thank you for considering my application.
</div>
<div>

Cheers,
Lei
</div>
        </article>

      </section>

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