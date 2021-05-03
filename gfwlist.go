package main

import (
	"regexp"
	"sort"
	"strings"

	"github.com/honwen/golibs/cip"
	"github.com/honwen/golibs/domain"
)

const officalGoogleDomain = `https://www.google.com/supported_domains`

var officalGFWListURLs = []string{
	`https://raw.githubusercontents.com/Loukky/gfwlist-by-loukky/master/gfwlist.txt`,
	`https://raw.githubusercontents.com/gfwlist/gfwlist/master/gfwlist.txt`,
}

const regxIP = `(25[0-5]|2[0-4]\d|[0-1]\d{2}|[1-9]?\d)\.(25[0-5]|2[0-4]\d|[0-1]\d{2}|[1-9]?\d)\.(25[0-5]|2[0-4]\d|[0-1]\d{2}|[1-9]?\d)\.(25[0-5]|2[0-4]\d|[0-1]\d{2}|[1-9]?\d)`

func filter(s []string, filterFunc func(string) bool) (after []string, fliterd []string) {
	for i := range s {
		if filterFunc(s[i]) {
			fliterd = append(fliterd, s[i])
		} else {
			after = append(after, s[i])
		}
	}
	return
}

func customsSort(domains []string, topDomains []string) []string {
	domains = domain.Sort(domains)
	regxIPv4 := regexp.MustCompile(cip.RegxIPv4)
	domains, ips := filter(domains, func(s string) bool {
		return regxIPv4.MatchString(s)
	})
	domains, extraTops := filter(domains, func(s string) bool {
		if strings.HasPrefix(s, "google.") || strings.HasPrefix(s, "blogspot.") {
			return true
		}
		i := sort.SearchStrings(topDomains, s)
		return (i < len(s) && domains[i] == s)
	})
	topDomains = domain.Sort(append(topDomains, extraTops...))

	sort.Strings(ips)
	sort.Strings(topDomains)
	domains = append(ips, append(topDomains, domains...)...)
	return domains
}

const initList = `
blogspot.com.by
blogspot.com.ee
blogspot.com.es
blogspot.in
blogspot.kr
blogspot.mx
blogspot.my
blogspot.pe
blogspot.qa
blogspot.sg
blogspot.ug
google.ac
google.berlin
google.cc
google.com.iq
google.com.lc
google.com.nf
google.com.tn
google.gf
google.gp
google.io
google.jp
google.net
google.ng
google.org
google.ph
google.sg
google.uk
google.ventures
appbridge.ca
bumptop.ca
mray.club
jav01.cc
jav168.cc
av1688.cc
avpanda.cc
javcc.cc
m-team.cc
hsxhr.cc
theav.cc
upjav.cc
airav.cc
avfox.cc
v.gd
womenwill.id
bazel.build
archive.md
zind.cloud
renzhe.cloud
edia.prod.mdn.mozit.cloud
async.be
weltweitwachsen.de
zukunftswerkstatt.de
clockwise.ee
lin.ee
g.page
mobileview.page
hanime1.me
download.91porn005.me
cslpldyb.me
calvappd.me
dayone.me
aibaobei.me
pexcn.me
cloudn.me
paofuyun.me
wuso.me
klip.me
conair.me
a.91gay.me
54647.online
ctyun.online
qingse.one
lululu.one
blinkload.zone
ping.pe
on.here
51dh.site
swag.live
n3ro.live
lih.kg
muncloud.dog
partylikeits1986.org
eveloper.mozilla.org
abpchina.org
hdrplusdata.org
coova.org
picasaweb.org
j2objc.org
cordcloud.org
openairinterface.org
iamremarkable.org
paxlicense.org
sciencemag.org
dartlang.org
verilystudywatch.org
brotli.org
meirifuli.org
ublock.org
blink.org
greasyfork.org
sleazyfork.org
edx-cdn.org
onefifteen.org
maven.org
dataliberation.org
yoyo.org
bumptop.org
cocoapods.org
cookiechoices.org
userstyles.org
digitalassetlinks.org
dn.mozillademos.org
bash-hackers.org
tt-rss.org
adblockplus.org
stxmosquitoproject.org
bitbucket.org
culturalspot.org
conscrypt.org
angulardart.org
openwrt.org
globaledu.org
download.i91av.org
w.org
mathjax.org
keytransparency.org
baselinestudy.org
studywatchbyverily.org
uProxy.org
getfoxyProxy.org
cnnmoney.ch
pageview.mobi
spoti.fi
ipn.li
hbogoasia.hk
page.link
googlecompare.co.uk
howtogetmo.co.uk
amazonaws.co.uk
thinkquarterly.co.uk
aaex.uk
x.team
overcast.fm
rime.im
area120.com
ohyeah1080.com
thtmod1.com
lhie1.com
agzy1.com
lsnzxzy1.com
gvt2.com
aimei133.com
xo104.com
sexzy4.com
office365.com
gvt5.com
gvt6.com
dkk37.com
sub147.com
jjdong7.com
gvt7.com
vgg8.com
icons8.com
cmail19.com
gvt9.com
comodoca.com
anaconda.com
tm.tnt-ea.com
porntea.com
blogsmithmedia.com
macromedia.com
0zmedia.com
wikipedia.com
hboasia.com
hbogoasia.com
txvia.com
ooyala.com
animezilla.com
garena.com
projectara.com
orbitera.com
picasa.com
velostrata.com
conviva.com
coova.com
symcb.com
p3.csgfnmdb.com
hindiweb.com
easybib.com
verilystudyhub.com
pubnub.com
tvb.com
encoretvb.com
9to5mac.com
cnbc.com
asp-cc.com
javcc.com
chroniclesec.com
cnnarabic.com
episodic.com
cloudmagic.com
tinypic.com
steamstatic.com
awsstatic.com
dropboxstatic.com
myfontastic.com
91.51rmc.com
megaupload.com
symcd.com
buzzfeed.com
snapseed.com
steampowered.com
admeld.com
googlefinland.com
javqd.com
smartmailcloud.com
leancloud.com
gmocloud.com
dlercloud.com
immxd.com
adobe.com
85tube.com
lubetube.com
tubetubetube.com
besthentaitube.com
pornstarbyface.com
squarespace.com
office.com
quickoffice.com
webappfieldguide.com
culturedcode.com
leetcode.com
mindnode.com
linode.com
api.amplitude.com
googlee.com
flirt4free.com
androidapksfree.com
fiftythree.com
avseesee.com
quiksee.com
isco.evergage.com
slack-edge.com
googlesverige.com
goolge.com
china-internet-exchange.com
theverge.com
nssurge.com
webpkgcache.com
toggleable.com
mashable.com
airtable.com
dribbble.com
netmarble.com
paddle.com
humblebundle.com
amytele.com
foofle.com
gogle.com
ggoogle.com
gooogle.com
froogle.com
jibemobile.com
novafile.com
gogole.com
keyhole.com
dnsimple.com
phobos.apple.com
tv.apple.com
name.com
fontawesome.com
projectbaseline.com
financeleadsonline.com
watchjavonline.com
javdoe.com
skype.com
webtype.com
920share.com
bumpshare.com
cryptocompare.com
tomshardware.com
hegre.com
nexitcore.com
imore.com
googlestore.com
thegooglestore.com
livefilestore.com
apture.com
azure.com
pstore.corpmerchandise.com
thecleversense.com
dn77.scoreuniverse.com
52hyse.com
pixate.com
gsuite.com
simplenote.com
onenote.com
evernote.com
beatthatquote.com
live.com
xboxlive.com
waze.com
getpricetag.com
foxdcg.com
yimg.com
yuetwanlauseng.com
bing.com
cryptographyengineering.com
ying.com
mdialog.com
blog.com
edgedatg.com
wzmyg.com
papalah.com
seselah.com
techcrunch.com
cobrasearch.com
cloudburstresearch.com
marketwatch.com
verilystudywatch.com
unsplash.com
pushwoosh.com
techsmith.com
symauth.com
vivaldi.com
mperial.insendi.com
postini.com
dazn-api.com
paddleapi.com
aerisapi.com
technorati.com
xiti.com
modmyi.com
evozi.com
littlehj.com
69vj.com
av6k.com
celestrak.com
slack.com
picnik.com
urofensk-prod-env.eu-west-1.elasticbeanstalk.com
uplynk.com
chromebook.com
gitbook.com
googledanmark.com
upwork.com
playstationnetwork.com
baicaonetwork.com
sonyentertainmentnetwork.com
andom.zendesk.com
coindesk.com
fanatical.com
lanternal.com
googlecapital.com
feedsportal.com
googel.com
crossmediapanel.com
screenwisetrendspanel.com
api.mixpanel.com
googl.com
mail.com
protonmail.com
fastmail.com
googil.com
womenwill.com
aol.com
trustasiassl.com
hulustream.com
cnnmoneystream.com
shazam.com
0emm.com
servebom.com
theplatform.com
crittercism.com
adobedtm.com
lithium.com
impermium.com
fanhaodian.com
follasian.com
typcn.com
netdna-cdn.com
vox-cdn.com
blogcdn.com
licdn.com
aolcdn.com
disquscdn.com
aspnetcdn.com
lightboxcdn.com
dazndn.com
manhuaren.com
globalsign.com
boomtrain.com
linkedin.com
logmein.com
avinin.com
ridepenguin.com
campuslondon.com
google-syndication.com
googleacquisitionmigration.com
playstation.com
pokemon.com
ichineseporn.com
myavfun.com
itgonglun.com
supermariorun.com
g-tun.com
avn.com
libsyn.com
dazn.com
indazn.com
nintendo.com
go.com
hbogo.com
hwgo.com
xteko.com
trello.com
xeeno.com
yahoo.com
usvimosquito.com
stcroixmosquito.com
stxmosquito.com
cdn.angruo.com
tdesktop.com
bumptop.com
getbumptop.com
alfredapp.com
getcloudapp.com
firebaseapp.com
ghostnoteapp.com
dueapp.com
pyhapp.com
ulyssesapp.com
ishowsapp.com
gipscorp.com
omnigroup.com
economistgroup.com
nutaq.com
mlssoccer.com
tinder.com
ees.elsevier.com
docker.com
schemer.com
hqporner.com
instapaper.com
mytvsuper.com
stc-server.com
pagespeedmobilizer.com
txmblr.com
googlr.com
droplr.com
useplannr.com
microsofttranslator.com
pixelmator.com
ycombinator.com
tablesgenerator.com
mfg-inspector.com
websnapr.com
crr.com
vultr.com
nianticlabs.com
cnnlabs.com
cnnpolitics.com
cloudrobotics.com
oss-cn-hongkong.aliyuncs.com
manyvids.com
screenwisetrends.com
digitaltrends.com
verilylifesciences.com
googlepages.com
oneworldmanystories.com
asianpornmovies.com
instructables.com
armorgames.com
dowjones.com
bumptunes.com
googleventures.com
ubertags.com
slack-msgs.com
cs4hs.com
arcgis.com
sharethis.com
yahooapis.com
moodstocks.com
parallels.com
kindgirls.com
mergersandinquisitions.com
bangbros.com
googlephotos.com
googlemaps.com
1ucrs.com
economistgroupcareers.com
macrumors.com
verizonwireless.com
ipaddress.com
ingress.com
flexibits.com
tapbots.com
api.termius.com
mocloudplus.com
foxplus.com
venturebeat.com
stripchat.com
brocaproject.com
debugproject.com
usvimosquitoproject.com
stcroixmosquitoproject.com
stxmosquitoproject.com
ampproject.com
deskconnect.com
engadget.com
wallet.com
pushbullet.com
cnet.com
helpshift.com
bandisoft.com
zuremarketplace.microsoft.com
zure.microsoft.com
eveloper.microsoft.com
bit.com
ifixit.com
ubnt.com
app-measurement.com
cdn.segment.com
jshint.com
codespot.com
conscrypt.com
pdfexpert.com
moves-export.com
fast.com
wtfast.com
nest.com
vaginacontest.com
todoist.com
javmost.com
stripst.com
usertrust.com
att.com
pittpatt.com
ifttt.com
orkut.com
sublimetext.com
eiu.com
jinnaju.com
heroku.com
r18lu.com
jiayoulu.com
721av.com
x99av.com
pigav.com
iijav.com
youav.com
buzzav.com
gv.com
api.tiktokv.com
revolv.com
gerritcodereview.com
dialogflow.com
hbonow.com
textnow.com
saynow.com
javynow.com
vzw.com
javhd3x.com
gdax.com
seqingx.com
japonx.com
japronx.com
box.com
xbox.com
fox.com
joox.com
solveforx.com
bkrtx.com
justgetflux.com
nekoxxx.com
bintray.com
mydirtyhobby.com
keytransparency.com
godaddy.com
baselinestudy.com
cnnmoney.com
typography.com
optimizely.com
sciencedaily.com
osxdaily.com
verily.com
studywatchbyverily.com
nexitally.com
thinkquarterly.com
sony.com
69story.com
fullstory.com
flurry.com
workflowy.com
freeopenProxy.com
4everProxy.com
publishproxy.com
oauthz.com
gstatic.cn
g.cn
gstaticcnapps.cn
gkecnapps.cn
pinboard.in
womenwill.in
press.vin
soirt4.fun
surge.run
archive.vn
vbstatic.co
macid.co
twimg.co
maying.co
scdn.co
amazon.co
vsco.co
insder.co
tmblr.co
tumblr.co
xplr.co
unblocksites.co
prfct.co
v2ex.co
wikileaks.info
xclient.info
abema.io
hayabusa.io
fabric.io
gsrc.io
openthread.io
appbridge.io
coinsquare.io
weverse.io
openweave.io
intercom.io
cnn.io
io.io
trouter.io
kfs.io
nomulus.foo
keytransparency.foo
javhd.pro
notion.so
ark.to
danilo.to
1jsa22.vip
1jjdg2.vip
0plkijj.vip
dvh30n.vip
japonx.vip
japronx.vip
cloudgarage.jp
shadowverse.jp
gosetsuden.jp
happyon.jp
go.jp
nhncorp.jp
cygames.jp
hulu.jp
localnetwork.uop
web.app
gridaware.app
run.app
bighead.group
womenwill.com.br
kat.cr
eurecom.fr
kakao.co.kr
lizzard.nefficient.co.kr
d.pr
itun.es
workflow.is
gfx.ms
sfx.ms
geni.us
invidio.us
disq.us
unfiltered.news
apple.news
sextop1.net
2o7.net
recaptcha.net
gabia.net
nteractive-examples.mdn.mozilla.net
coova.net
picasaweb.net
omtrdc.net
yastatic.net
sstatic.net
nsstatic.net
timeinc.net
gaypad.net
akamaized.net
txcloud.net
llnwd.net
office.net
javhdfree.net
azureedge.net
akamaiedge.net
msedge.net
adgoogle.net
edgesuite.net
live.net
pximg.net
bing.net
dartsearch.net
akamai.net
js.revsci.net
vivaldi.net
gonglchuangl.net
javfull.net
fzlm.net
piaotian.net
line-cdn.net
kakaocdn.net
2mdn.net
onefifteen.net
johren.net
playstation.net
osakamotion.net
tokyomotion.net
nintendo.net
n3ro.net
bumptop.net
cloudapp.net
hockeyapp.net
tbr.tangbr.net
tweetmarker.net
ooklaserver.net
japanesebeauties.net
apple-dns.net
akadns.net
tapbots.net
south-plus.net
stxmosquitoproject.net
ampproject.net
hinet.net
typekit.net
speedsmart.net
googlecert.net
cdnst.net
entrust.net
reimu.net
demdex.net
kphimsex.net
rightcove.imgix.net
japonx.net
japronx.net
box.net
edgekey.net
cachefly.net
fastly.net
appbridge.it
cnn.it
gfw.report
tssp.best
avmoo.cyou
jav.guru
stadia.dev
fuchsia.dev
g.dev
bdn.dev
chromeos.dev
gateway.dev
osha.gov
nasa.gov
ssa.gov
state.gov
nih.gov
uspto.gov
usgs.gov
census.gov
avsee01.tv
av01.tv
520aa.tv
abema.tv
a.kslive.tv
7mm.tv
fubo.tv
trakt.tv
hpjav.tv
popjav.tv
japonx.tv
japronx.tv
meet.new
devsitetest.how
lp99.pw
ssplive.pw
ssrpass.pw
st.luluku.pw
isexomega.tw
isexlove.tw
kingkong.com.tw
kkbox.com.tw
linetv.tw
womenwill.mx
near.by
cl.ly
simp.ly
chronicle.security
love7.xyz
tkb008.xyz
jgg18.xyz
javdove8.xyz
caime.xyz
mdlf.xyz
gouri.xyz
18novel.xyz
limbopro.xyz
aavs.xyz
114av.xyz
16fhgdty.xyz
`
