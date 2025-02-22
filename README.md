# webscraper

github repository: https://github.com/mwood881/webscraper.git

## Project Description:
web scraping text from Wikipedia pages and converting them into a JSON lines file also known as a Wikipedia Scraper. 

## Features:
1) There are a list of URLs in main.go that was given in the assignment.
2) Uses concurrent scraping with goroutines.
3) the main text is scrape from Wikipedia urls.
4) The data is stored in scraped_output.jl
5) the time for the program to run is also reported.

## Prerequisites:
Install the following:
1. Go https://go.dev/doc/install
2. colly package (install it via go get -u github.com/gocolly/colly)


## Get Started:
1. Clone the repository
```bash
git clone https://github.com/mwood8881/webscraper.git
cd webscraper

```

2) Install colly:
```bash
go get -u github.com/gocolly/colly

```

3) build the program which is shown in main.go

4) Run the program:
   
5) ```bash
go run main.go
```

6) output file is in scraped_data.jl

## Code Example

url: is the url of the data
text: is the scraped text

{"url":"https://en.wikipedia.org/wiki/Applications_of_artificial_intelligence","text":"\n\n.mw-parser-output .ambox{border:1px solid #a2a9b1;border-left:10px solid #36c;background-color:#fbfbfb;box-sizing:border-box}.mw-parser-output .ambox+link+.ambox,.mw-parser-output .ambox+link+style+.ambox,.mw-parser-output .ambox+link+link+.ambox,.mw-parser-output .ambox+.mw-empty-elt+link+.ambox,.mw-parser-output .ambox+.mw-empty-elt+link+style+.ambox,.mw-parser-output .ambox+.mw-empty-elt+link+link+.ambox{margin-top:-1px}html body.mediawiki .mw-parser-output .ambox.mbox-small-left{margin:4px 1em 4px 0;overflow:hidden;width:238px;border-collapse:collapse;font-size:88%;line-height:1.25em}.mw-parser-output .ambox-speedy{border-left:10px solid #b32424;background-color:#fee7e6}.mw-parser-output .ambox-delete{border-left:10px solid #b32424}.mw-parser-output .ambox-content{border-left:10px solid #f28500}.mw-parser-output .ambox-style{border-left:10px solid #fc3}.mw-parser-output .ambox-move{border-left:10px solid #9932cc}.mw-parser-output .ambox-protection{border-left:10px solid #a2a9b1}.mw-parser-output .ambox .mbox-text{border:none;padding:0.25em 0.5em;width:100%}.mw-parser-output .ambox .mbox-image{border:none;padding:2px 0 2px 0.5em;text-align:center}.mw-parser-output .ambox .mbox-imageright{border:none;padding:2px 0.5em 2px 0;text-align:center}.mw-parser-output .ambox .mbox-empty-cell{border:none;padding:0;width:1px}.mw-parser-output .ambox .mbox-image-div{width:52px}@media(min-width:720px){.mw-parser-output .ambox{margin:0 10%}}@media print{body.ns-0 .mw-parser-output .ambox{display:none!important}}.mw-parser-output .multiple-issues-text{width:95%;margin:0.2em 0}.mw-parser-output .multiple-issues-text\u003e.mw-collapsible-content{margin-top:0.3em}.mw-parser-output .compact-ambox .ambox{border:none;border-collapse:collapse;background-color:transparent;margin:0 0 0 1.6em!important;padding:0!important;width:auto;display:block}body.mediawiki .mw-parser-output .compact-ambox .ambox.mbox-small-left{font-size:100%;width:auto;margin:0}.mw-parser-output .compact-ambox .ambox .mbox-text{padding:0!important;margin:0!important}.mw-parser-output .compact-ambox .ambox .mbox-text-span{display:list-item;line-height:1.5em;list-style-type:disc}body.skin-minerva .mw-parser-output .multiple-issues-text\u003e.mw-collapsible-toggle,.mw-parser-output .compact-ambox .ambox .mbox-image,.mw-parser-output .compact-ambox .ambox .mbox-imageright,.mw-parser-output .compact-ambox .ambox .mbox-empty-cell,.mw-parser-output .compact-ambox .hide-when-compact{display:none}This article has multiple issues. Please help improve it or discuss these issues on the talk page. (Learn how and when to remove these messages)\n\n      This article's tone or style may not reflect the encyclopedic tone used on Wikipedia. See Wikipedia's guide to writing better articles for suggestions.  (April 2022) (Learn how and when to remove this message)\nThis article may lend undue weight to very obscure AI projects of questionable importance. Please help improve it by rewriting it in a balanced fashion that contextualizes different points of view.  (April 2022) (Learn how and when to remove this message)\n    \n (Learn how and when to remove this message)\n.mw-parser-output .hlist dl,.mw-parser-output .hlist ol,.mw-parser-output .hlist ul{margin:0;padding:0}.mw-parser-output .hlist dd,.mw-parser-output .hlist dt,.mw-parser-output .hlist li{margin:0;display:inline}.mw-parser-output .hlist.inline,.mw-parser-output .hlist.inline dl,.mw-parser-output .hlist.inline ol,.mw-parser-output .hlist.inline ul,.mw-parser-output .hlist dl dl,.mw-parser-output .hlist dl ol,.mw-parser-output .hlist dl ul,.mw-parser-output .hlist ol dl,.mw-parser-output .hlist ol ol,.mw-parser-output .hlist ol ul,.mw-parser-output .hlist ul dl,.mw-parser-output .hlist ul ol,.mw-parser-output .hlist ul ul{display:inline}.mw-parser-output .hlist .mw-empty-li{display:none}.mw-parser-output .hlist dt::after{content:\": \"}.mw-parser-output .hlist dd::after,.mw-parser-output .hlist li::after{content:\" · \";font-weight:bold}.mw-parser-output .hlist dd:last-child::after,.mw-parser-output .hlist dt:last-child::after,.mw-parser-output .hlist li:last-child::after{content:none}.mw-parser-output .hlist dd dd:first-child::before,.mw-parser-output .hlist dd dt:first-child::before,.mw-parser-output .hlist dd li:first-child::before,.mw-parser-output .hlist dt dd:first-child::before,.mw-parser-output .hlist dt dt:first-child::before,.mw-parser-output .hlist dt li:first-child::before,.mw-parser-output .hlist li dd:first-child::before,.mw-parser-output .hlist li dt:first-child::before,.mw-parser-output .hlist li li:first-child::before{content:\" (\";font-weight:normal}.mw-parser-output .hlist dd dd:last-child::after,.mw-parser-output .hlist dd dt:last-child::after,.mw-parser-output .hlist dd li:last-child::after,.mw-parser-output .hlist dt dd:last-child::after,.mw-parser-output .hlist dt dt:last-child::after,.mw-parser-output .hlist dt li:last-child::after,.mw-parser-output .hlist li dd:last-child::after,.mw-parser-output .hlist li dt:last-child::after,.mw-parser-output .hlist li li:last-child::after{content:\")\";font-weight:normal}.mw-parser-output .hlist ol{counter-reset:listitem}.mw-parser-output .hlist ol\u003eli{counter-increment:listitem}.mw-parser-output .hlist ol\u003eli::before{content:\" \"counter(listitem)\"\\a0 \"}.mw-parser-output .hlist dd ol\u003eli:first-child::before,.mw-parser-output .hlist dt ol\u003eli:first-child::before,.mw-parser-output .hlist li ol\u003eli:first-child::before{content:\" (\"counter(listitem)\"\\a0 \"}.mw-parser-output .sidebar{width:22em;float:right;clear:right;margin:0.5em 0 1em 1em;background:var(--background-color-neutral-subtle,#f8f9fa);border:1px solid var(--border-color-base,#a2a9b1);padding:0.2em;text-align:center;line-height:1.4em;font-size:88%;border-collapse:collapse;display:table}body.skin-minerva .mw-parser-output .sidebar{display:table!important;float:right!important;margin:0.5em 0 1em 1em!important}.mw-parser-output .sidebar-subgroup{width:100%;margin:0;border-spacing:0}.mw-parser-output .sidebar-left{float:left;clear:left;margin:0.5em 1em 1em 0}.mw-parser-output .sidebar-none{float:none;clear:both;margin:0.5em 1em 1em 0}.mw-parser-output .sidebar-outer-title{padding:0 0.4em 0.2em;font-size:125%;line-height:1.2em;font-weight:bold}.mw-parser-output .sidebar-top-image{padding:0.4em}.mw-parser-output .sidebar-top-caption,.mw-parser-output .sidebar-pretitle-with-top-image,.mw-parser-output .sidebar-caption{padding:0.2em 0.4em 0;line-height:1.2em}.mw-parser-output .sidebar-pretitle{padding:0.4em 0.4em 0;line-height:1.2em}.mw-parser-output .sidebar-title,.mw-parser-output .sidebar-title-with-pretitle{padding:0.2em 0.8em;font-size:145%;line-height:1.2em}.mw-parser-output .sidebar-title-with-pretitle{padding:0.1em 0.4em}.mw-parser-output .sidebar-image{padding:0.2em 0.4em 0.4em}.mw-parser-output .sidebar-heading{padding:0.1em 0.4em}.mw-parser-output .sidebar-content{padding:0 0.5em 0.4em}.mw-parser-output .sidebar-content-with-subgroup{padding:0.1em 0.4em 0.2em}.mw-parser-output .sidebar-above,.mw-parser-output .sidebar-below{padding:0.3em 0.8em;font-weight:bold}.mw-parser-output .sidebar-collapse .sidebar-above,.mw-parser-output .sidebar-collapse .sidebar-below{border-top:1px solid #aaa;border-bottom:1px solid #aaa}.mw-parser-output .sidebar-navbar{text-align:right;font-size:115%;padding:0 0.4em 0.4em}.mw-parser-output .sidebar-list-title{padding:0 0.4em;text-align:left;font-weight:bold;line-height:1.6em;font-size:105%}.mw-parser-output .sidebar-list-title-c{padding:0 0.4em;text-align:center;margin:0 3.3em}@media(max-width:640px){body.mediawiki .mw-parser-output .sidebar{width:100%!important;clear:both;float:none!important;margin-left:0!important;margin-right:0!important}}body.skin--responsive .mw-parser-output .sidebar a\u003eimg{max-width:none!important}@media screen{html.skin-theme-clientpref-night .mw-parser-output .sidebar:not(.notheme) .sidebar-list-title,html.skin-theme-clientpref-night .mw-parser-output .sidebar:not(.notheme) .sidebar-title-with-pretitle{background:transparent!important}html.skin-theme-clientpref-night .mw-parser-output .sidebar:not(.notheme) .sidebar-title-with-pretitle a{color:var(--color-progressive)!important}}@media screen and (prefers-color-scheme:dark){html.skin-theme-clientpref-os .mw-parser-output .sidebar:not(.notheme) .sidebar-list-title,html.skin-theme-clientpref-os .mw-parser-output .sidebar:not(.notheme) .sidebar-title-with-pretitle{background:transparent!important}html.skin-theme-clientpref-os .mw-parser-output .sidebar:not(.notheme) .sidebar-title-with-pretitle a{color:var(--color-progressive)!important}}@media print{body.ns-0 .mw-parser-output .sidebar{display:none!important}}Part of a series onArtificial intelligence (AI)\nMajor goals\nArtificial general intelligence\nIntelligent agent\nRecursive self-improvement\nPlanning\nComputer vision\nGeneral game playing\nKnowledge reasoning\nNatural language processing\nRobotics\nAI safety\n\nApproaches\nMachine learning\nSymbolic\nDeep learning\nBayesian networks\nEvolutionary algorithms\nHybrid intelligent systems\nSystems integration\n\nApplications\nBioinformatics\nDeepfake\nEarth sciences\n Finance \nGenerative AI\nArt\nAudio\nMusic\nGovernment\nHealthcare\nMental health\nIndustry\nTranslation\n Military \nPhysics\nProjects\n\nPhilosophy\nArtificial consciousness\nChinese room\nFriendly AI\nControl problem/Takeover\nEthics\nExistential risk\nTuring test\nUncanny valley\n\nHistory\nTimeline\nProgress\nAI winter\nAI boom\n\nGlossary\nGlossary\n.mw-parser-output .navbar{display:inline;font-size:88%;font-weight:normal}.mw-parser-output .navbar-collapse{float:left;text-align:left}.mw-parser-output .navbar-boxtext{word-spacing:0}.mw-parser-output .navbar ul{display:inline-block;white-space:nowrap;line-height:inherit}.mw-parser-output .navbar-brackets::before{margin-right:-0.125em;content:\"[ \"}.mw-parser-output .navbar-brackets::after{margin-left:-0.125em;content:\" ]\"}.mw-parser-output .navbar li{word-spacing:-0.125em}.mw-parser-output .navbar a\u003espan,.mw-parser-output .navbar a\u003eabbr{text-decoration:inherit}.mw-parser-output .navbar-mini abbr{font-variant:small-caps;border-bottom:none;text-decoration:none;cursor:inherit}.mw-parser-output .navbar-ct-full{font-size:114%;margin:0 7em}.mw-parser-output .navbar-ct-mini{font-size:114%;margin:0 4em}html.skin-theme-clientpref-night .mw-parser-output .navbar li a abbr{color:var(--color-base)!important}@media(prefers-color-scheme:dark){html.skin-theme-clientpref-os .mw-parser-output .navbar li a abbr{color:var(--color-base)!important}}@media print{.mw-parser-output .navbar{display:none!important}}vte\nArtificial intelligence (AI) has been used in applications throughout industry and academia. In a manner analogous to electricity or computers, AI serves as a general-purpose technology. AI programs are designed to simulate human perception and understanding. These systems are capable of adapting to new information and responding to changing situations. Machine learning has been used for various scientific and commercial purposes[1] including language translation, image recognition, decision-making,[2][3] credit scoring, and e-commerce.\n\n\nInternet and e-commerce[edit]\n.mw-parser-output .hatnote{font-style:italic}.mw-parser-output div.hatnote{padding-left:1.6em;margin-bottom:0.5em}.mw-parser-output .hatnote i{font-style:normal}.mw-parser-output .hatnote+link+.hatnote{margin-top:-0.5em}@media print{body.ns-0 .mw-parser-output .hatnote{display:none!important}}Main article: Marketing and artificial intelligence\nWeb feeds and posts[edit]\nMachine learning has been used for recommendation systems in determining ..."

## No tests needed according to canvas page

## AI Help
I utilized AI to assist in creating code as shown in the AI Helper file
I also used the github copilot at the end by asking " make this code better" and utilized the code it gave me for more efficient coding formats. 

