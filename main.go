package main

import (
	"github.com/psyb0t/go-spamcheck/spamcheck"
	"fmt"
	"encoding/json"
)

func main() {
	// just a random email message i got on https://www.guerrillamail.com/ using a@grr.la
	rawEmail := `Delivered-To: a@sharklasers.com
Received: from mta03.xxxpersonals.com (mta03.xxxpersonals.com  [52.22.58.118])
	by grr.la with SMTP id b0de3b3e98a76136473fbc17cadbf6df@grr.la;
	Thu, 24 Jan 2019 22:05:51 +0000
DKIM-Signature: v=1;a=rsa-sha1;s=dkim;d=xxxpersonals.com;c=simple/simple;q=dns/txt;l=5;bh=Bhpq/Jg2ftVZ7zYadMR0/2Yp7zM=;h=X-Track-Data:Message-ID:From:List-Unsubscribe;b=acmz8FWnhgALRTCOcsoHZRWsgnOPsTpQVqnPS0d5FSK0ZBsM0Ck0OAHWvuGheUg04BXrI5foNVEH7yZo51Xk2+8SlF8IYF7EBgKBnqxrCDo6lHBioliYOlUSayYbsk2qUmWGu14nsIpLFtTVPyeM8r63GRR31LVM/visc0qFE0w=
X-Track-Data: 5w6l81kx8pk8j77mc0z5v5t0d7jdbphAtw6dr592k47s5ym0nngqw38
Sender: <bounce@link.xxxpersonals.com>
Message-ID: <5w6l81kx8pk8j77mc0z5v5t0d7jdbphAtw6dr592k47s5ym0nngqw38@link.xxxpersonals.com>
X-Data: 1vM3dWCjGbEifRrY5tO23DInV3O9UWjMnF0VRCCLEJPxO0Eez68TSSDuPsnP1_nttsOL60k4sPlqRkqos61TYUtNOL64ymzMCJZqC8UyZULWRGLQTiWJfzal9NNtGzVGj!HdaJdTO3SMqmcBNUnBplS1sJcwQCJs868NW4ixbX1itO9uy9bno6qFOMENbWqpK5Z3NJ4frYFJMmGX!HVdsXkI7ml0bVEIeWTj!MjZSd3SNNzYyXjEV2r9!N3xdZqPdbyK4NAj2lePCjfJA8TKoErBVJbxQ37p6YPS8wHkc7BtXwZ!BLkRgOBKYeDVWi7iqSiEqOXW1h645ONutxhCk4lldQoQERZYamEAAi1vSDd4z!ynuzC8H4JLm8eT!8z53KVzv!iqf4pfFMbJ8KllL!Un7jkxin!QSE1pdnUC5HEMUA75MtvJ0ShtrKDGjCDbYaNlMre9P_XheiuliSPAkvxJSRm_XJ9j5Vs5vp8Ay!HLqCR2KwheDROlINWXmiFbJ28JJ_Zmd8Vz705Iwf8t4HWKA9WElNa211pQ64QJu9CefmYBHlP6Cfuc80psZKqaZK1tpI6kEHnQNGrFb!tQLhR1r!zbvH2lawnWEuEO4_EAjppJFZXSiU2As_dHR1Wf0fe4!Pcq7z21f!x5kqx3GHw2UpJ9aniZKL3zKomUllU
From: xxxpersonals <support@xxxpersonals.com>
To: a@grr.la
Subject: Jenna sent you a sexy message 
List-Unsubscribe: <http://link.xxxpersonals.com/qa/c6rdqhdv54mbrm91hmr8Aj5p183z1wfA8pdznb3rs6vlnl9l13nAmqvr88y4b787rpw1sz7yzA21xfymx6s9pg4jnj4py6pwdtm89h935q1l6q41k3p3mp8cAcvz68mjnt3wyn6kdf03krp6l41v8557tq2ct2qjmgp3vqbqfx1khs8kr9dqlygbf4lgd57vn3f9p3yclzmwc8xy819gfkdq2yg6jx416mm23fA34mjt7hz77rnb88wf4s6zd7vwgps9dxlfllrqv5fqhv92fjy0vjpxc5cq0vf3hrvfkk40v5Al5bf7223234y18rfs61y0vAdq30qqygvA9lggb3pAvk6A3p8r898b1k56qky7ck7qlty5vtj7tk8d5217xhw7ssj5hk4hcrwzlvcA9dq8pyp2w88fbmAdchxgsfs6mg9x3lj5x91cz4ymkm93j9f3tpllkcdA682cn1xk1lj5t6px07rzlzm3h4thnfAAsg6s3bp396hh405h4jflmrycmrmpj2hj7nmhrzhAtsm5rnwjxc5d4pj46k63jn560l5s5kbg0gwz79kwcyvv4m24cp9q2k8hh98fx5gl33k240yh9Atb9nk189rw98hvctxjn60fgt5sgqqp4ry5sm7kxz3498wtqsqpgpdrAk4vq8jlnxlt014qkngqp7sm21xrcm8h26l3rfrAfchrrs9z2vcbArcqlzqmfshsrg70Abntp3wqkzc99Amxtnn7394l5f3rz8r5nf2dm3wj0sddjAn7lkkj2Aj93jm91x7hnrp4vdn1vhv92b8ntq6g0821lg3vpt6q0rpvpq1fx9040sx7kzznlbtt6f0md6b20dn836d84ftsqs18gf79nyA0pf85v4p64vd2sgmqb3rdf4yq2220zbxzhy2lnz811wmwj1hl5gp8l2wqwktdsykcqp2cqydydtzsgt7d6yfAp3wsg8z6wfzt957p744hfm261g2cfsmcxAp0kqjjbz613rq85ntwt8zksggm3j4j17fmcpyAygv7xrzk7192gtkt23dwl3fA1sqv965lkll0xwnr91A65bfd2gnqA4jv035m8vlqdph5rkj7nwv6j2tnhjm1c92scsmbykrhr6Azvlvhck8Al1zbkdfgbsrb3f9j74d75A2z25tpg03Ajkr2shcq5v410A11pv19qy5mqhn3mwyt8xjnzsxnp1mqd2hps0n6ttlsg1d5t74r2wv7kknfj7ss5hv4vpjbbzAtv1xf24A2rl01wsbp5v7d82qxchbn2kb5pq5y6A8jkn0hsmk2xwl1ll2AzcbAz9z1tv2vqzAA6zj38xtsp1rvjr0fv42006d586wxc2hq0h0tw4nwfzhy7wmrpt0vssfm5clvcp9grk10th9gfy372szAbcrpjm7kmm2hcwnA17wqff9gmsAd8mn2A0dj88wwnjv9l51k03z823fkb7mw4vkd067ny48xcvj8jkcsg3sxxl8pfzjsjmbxp5k8lcttydnj3A77r52xkqAv/k38>,<mailto:listunsubscribe@xxxpersonals.com>
Content-Type: multipart/alternative;
 boundary="=_4fe15f765d8e68d30d2f92923456461f"
MIME-Version: 1.0
X-rpcampaign: stallionfsc382063190124t20190124050551

--=_4fe15f765d8e68d30d2f92923456461f
Content-Transfer-Encoding: quoted-printable
Content-Type: text/html

<!DOCTYPE html PUBLIC "-//W3C//DTD HTML 4.01 Transitional//EN" "http://www=
=2Ew3.org/TR/html4/loose.dtd">

<html>
  <head>
    <meta http-equiv=3D"Content-Type" content=3D"text/html; charset=3Dutf-8=
">
    <title>Jenna sent you a sexy message =F0=9F=92=AC</title>

  </head>
  <body style=3D"background: linear-gradient(to right, #ffffff, #e6ffff); f=
ont-family: Verdana, sans-serif; color: #880000">
    <!-- Template 3a - (04) -->
    <span class=3D"preheader" style=3D"display: none !important; visibility=
: hidden; opacity: 0; color: transparent; height: 0; width: 0;">
      Jenna sent you a sexy message =F0=9F=92=AC
    </span>
    <center>
      <table cellpadding=3D"0" cellspacing=3D"0" style=3D"width: 100%;">
        <tbody>
          <tr>

            <td align=3D"right">
              <a href=3D"http://link.xxxpersonals.com/sx/9dfhr4cs2jpxy5sy4j=
p38mmrrk77Ay2ptlsA5n5l9vthzs5tsrb5815grcslgy7xhkmvvzymyb6qkmc65gr017bc7wgdq=
m67qx30rx2knb2d341t24frkc1y7kmmkrsnypz44y4ry49vA0zbfy3pt24t8j56pvbcygjbht8c=
539bq8w00tnb0f5pfcf3n7twmyA92xAxslchqgwf8x78kvkAkA9hys0phg9b29shq3zz51sskgj=
qlffmm7zhrdx7rlkqhs6y18prh587r3mkdr53xx8hfm7fpkvv3r0pfj8zrlwbAgdjxgrw0fj25r=
p484klvw72hvwj4xfx1hm80y44122jlgtdg4mn1fy6b7973fc99wyd0x5kc200d7hw30zA1vl9p=
Ahn3689z0ffm28dy0q1z4f4g7vfxny0s3ktvAj8Asnlr9grfd3c1f6b64jg0dnw75yhsrmjn5vh=
gc30wx33nl2nq99zdzmb0tp6xc0gt2sjh5fcvt55sq508ls84ghmA6r3dsmrlxvq4ykdjc0l4ns=
4A7fykrt6t1g8mwv3znjA5A1hArryjmqr5qb8788tkjznw9mjzfb6sbclyk04ybby8kh2mh34nc=
w5vgc95AybgArcAy00tr6xrgm50hm9ts4jh67sAr5prq1xqtvwvktqt6dsjvj74bgnqsksrAqs1=
81g9b1xwhxzrA84yr7hb28hhmw6v2qxqhmxwlc66rnmt1m81byd72pjs0zf2qq8j6fj6fwv3l28=
xyA5smgygf7c04n79Ay5526vykf66bzAy95k326d9ctwvptps31dcy868cjAg2twnb53rlq3Akf=
3rcn9zs0kfgg7tcfrb1hkb8rh86l9hdh6xnfzdd52Ac7jwk4hfkvrgyh30ylkjr29761ykdA1s2=
vqmvwgf3986v0k7mbAbj11q27h7c0r7Axxd36ppbn4bwdvkj3rwyztjhcsdq5q4gzt8jk1txv1t=
kbjqdjh44y3p0jmlwb41lw9wkgs3zn8j074n33wxv49ydl0dn1hn602qx8dc1bsmgqp790ddkbd=
cfd2pnsbh5p2rx1671fbd63bh3s09djtA82q7dvdms2h4rkxgdxA9w0g3bdbrnlvmrrt7bkcddr=
A9p8p16cyynphqmhpfhl9rbxm1hrcxAjApwpnm2qnrtc3cnxfqtqyjxvs5fAln10tfvt19kzqlq=
xy4k83cxbn8hlA4n2A9btvpxgw62w30f3gwklj9h1t03jAzdkjAkc6ck7cyxk1159z5m2rny34t=
gwbzgq9cp16hhr2qpsl3twfxzqdp66fwp1zmqtrc6mtzqxA6xttq1qlh4frn8d2dp8tkcjn4qvs=
jz7bxvxn222fkzbcn6bqnbmt0tmvk5xAyq9gdwpd01t00ghzlA3n6t5kAqql5krs31hgjn3f1qb=
w8yd74mcmbd1gA8zvh6lm7z4kgxn4lft0hxhzbsq9whApp4mvr4rldzAntjpglhvwyq9bq0lvw0=
5lp34pfgrv0jyjc8bhqryqz9b3zk41q5pp4cAwg5n393/d38"></a>
            </td>
          </tr>
        </tbody>
      </table>
      <table cellpadding=3D"0" cellspacing=3D"0" style=3D"width: 100%;">
        <tbody>
          <tr>
            <td align=3D"left">
              <br>&nbsp;Dear A1999xmnds,
              <br>
              <br>&nbsp;Jenna sent you a sexy message
              <br>
              <br>&nbsp;<a href=3D"http://link.xxxpersonals.com/rg/qv6mh9Av=
qsdnc7bk6hfkqzzdtlgzvnsx1xA7bsgmAv098p69b2vl566w0pwmy4flz6n0fAjtvcfxt85rr3h=
fbn3bncx2k1899lpxymwb3jyyz7qtb71y8Av6rxhz5v94k1scddAcr553080b4y7f26md5c3z6q=
c606mwfpmxwcvrzcd69w31cpbz5kwcl1h45ncp4q0lfddltm5j4hxqs2p18gw8mgtkftdy61qk9=
ktpy9q1fz1dAA88930frAc3sxxjtf1pm85wk0vg86h753ypwfj00f3skt2vgtlf91n4jA0yypn8=
4b75k09nfhtkrnrnhkvfp453mtks0bmvn5mvfq0804fnc4l5h1qt7g0tjbbgxmr0kclmtv28m8y=
nrz5cxnv32lb7zfttzcq2tsxchgstc41cjpc87f4vmy5Ak4Awk5knclrz7v9sq0tq79z23lzpww=
2bk82v54bgfyj9728gyw50csk0kAzqf72ms18jgvmz20hc46c872qsl85ll63kbkxrwvlnvwqrr=
dA5pcd3qA02c7m8tnf5Args22Aq4pndlsnm5ymvflghsbgqtrrjjxwn25km7mn8kz5grtgftvmf=
16w91Axztgcwfnt6nk6jhAh9z52wkjdmp7tlzq9f0k2ws3y9ggt3vd7r36gl31lch7k9t8487l8=
l3hclfmpvgpw8kk16ktf2mA1yphg7jth9bc6gz2q98sp39jvvyyb8hlfw47htskrk10686z5lnk=
trcqqsrz7gtA08pzvtc1kqj2hj88mw5xc6qggcgqbnnpqgb8127m67vy6cs5gb46z1kc0qzmf2z=
p5xz1ynj41v9b7rkvvrdjzzkyr7tjfbc3zc747xgw7x95n1mqvczqvqkym6mbglt3b48A6169tj=
1w8f9x9ptbrqnzcbxqv1g051g5nA877fpr0vlrykqwsb2twyvmfl5lz22bndghj2xg096fmwv5r=
x7z3qs16frjlf02dx8zdnlhk5hkyw13bwtw7f6j74wA2cv07wAxxj6mqwrmjs6k9df8t98nyl1t=
6ly8ghyz4hAA35kpr6qzzwc91w310dz3s95l87br6rqkdl9rbjf379A80y4pAjg6js8h0xv3bgb=
5147Abd846pd326yAf07156dq04hb6Af4yt6rdtyy9yn9n5tl6941gr3842ty6wb5tv0fy9w1vs=
n3rs4Ax7sqc3stwhc9bAAlvskym4s9g0tbk23cfA4nll26slg6wcfmAxhn0tdbvhqr9wmfkc9xg=
42lf7slp39xgrjd90mpA99xs58kpy14tt1h2sdrp5w7mj3nzhqvfAb76g9jlxy9dxwdwymsbrxA=
g500mr9ms5splclx0bndy5psntsp62tc9dn6k2fd6ppxxpAd6pbht7wtyql0c18qjc9q7wdmj53=
5gxmy95t2cdjhw1wn3qxgfv5lAh8hx47t0c90jmv9nzlm861tvt5/a38">Click Here To Rea=
d It</a>
              <br>
            </td>
          </tr>
        </tbody>
      </table>
      <table width=3D"100%" border=3D"0" cellspacing=3D"0" cellpadding=3D"0=
">
        <tbody>
          <tr>
            <td height=3D"20"></td>
          </tr>
          <tr>
            <td style=3D"text-align:center;  font-size:11px; color:#999; pa=
dding-bottom:15px;">
              <hr>
              You got this message because you signed up to our site. If yo=
u have any questions/comments regarding the site, please view our Terms or =
<a href=3D"mailto:support@xxxpersonals.com" target=3D"_blank">contact us</a=
>. If you no longer would like to receive messages from us please <a href=
=3D"http://link.xxxpersonals.com/sx/9dfhr4cs2jpxy5sy4jp38mmrrk77Ay2ptlsA5n5=
l9vthzs5tsrb5815grcslgy7xhkmvvzymyb6qkmc65gr017bc7wgdqm67qx30rx2knb2d341t24=
frkc1y7kmmkrsnypz44y4ry49vA0zbfy3pt24t8j56pvbcygjbht8c539bq8w00tnb0f5pfcf3n=
7twmyA92xAxslchqgwf8x78kvkAkA9hys0phg9b29shq3zz51sskgjqlffmm7zhrdx7rlkqhs6y=
18prh587r3mkdr53xx8hfm7fpkvv3r0pfj8zrlwbAgdjxgrw0fj25rp484klvw72hvwj4xfx1hm=
80y44122jlgtdg4mn1fy6b7973fc99wyd0x5kc200d7hw30zA1vl9pAhn3689z0ffm28dy0q1z4=
f4g7vfxny0s3ktvAj8Asnlr9grfd3c1f6b64jg0dnw75yhsrmjn5vhgc30wx33nl2nq99zdzmb0=
tp6xc0gt2sjh5fcvt55sq508ls84ghmA6r3dsmrlxvq4ykdjc0l4ns4A7fykrt6t1g8mwv3znjA=
5A1hArryjmqr5qb8788tkjznw9mjzfb6sbclyk04ybby8kh2mh34ncw5vgc95AybgArcAy00tr6=
xrgm50hm9ts4jh67sAr5prq1xqtvwvktqt6dsjvj74bgnqsksrAqs181g9b1xwhxzrA84yr7hb2=
8hhmw6v2qxqhmxwlc66rnmt1m81byd72pjs0zf2qq8j6fj6fwv3l28xyA5smgygf7c04n79Ay55=
26vykf66bzAy95k326d9ctwvptps31dcy868cjAg2twnb53rlq3Akf3rcn9zs0kfgg7tcfrb1hk=
b8rh86l9hdh6xnfzdd52Ac7jwk4hfkvrgyh30ylkjr29761ykdA1s2vqmvwgf3986v0k7mbAbj1=
1q27h7c0r7Axxd36ppbn4bwdvkj3rwyztjhcsdq5q4gzt8jk1txv1tkbjqdjh44y3p0jmlwb41l=
w9wkgs3zn8j074n33wxv49ydl0dn1hn602qx8dc1bsmgqp790ddkbdcfd2pnsbh5p2rx1671fbd=
63bh3s09djtA82q7dvdms2h4rkxgdxA9w0g3bdbrnlvmrrt7bkcddrA9p8p16cyynphqmhpfhl9=
rbxm1hrcxAjApwpnm2qnrtc3cnxfqtqyjxvs5fAln10tfvt19kzqlqxy4k83cxbn8hlA4n2A9bt=
vpxgw62w30f3gwklj9h1t03jAzdkjAkc6ck7cyxk1159z5m2rny34tgwbzgq9cp16hhr2qpsl3t=
wfxzqdp66fwp1zmqtrc6mtzqxA6xttq1qlh4frn8d2dp8tkcjn4qvsjz7bxvxn222fkzbcn6bqn=
bmt0tmvk5xAyq9gdwpd01t00ghzlA3n6t5kAqql5krs31hgjn3f1qbw8yd74mcmbd1gA8zvh6lm=
7z4kgxn4lft0hxhzbsq9whApp4mvr4rldzAntjpglhvwyq9bq0lvw05lp34pfgrv0jyjc8bhqry=
qz9b3zk41q5pp4cAwg5n393/d38">opt-out</a>.

            </td>
          </tr>
        </tbody>
      </table>
      <table cellpadding=3D"0" cellspacing=3D"0" style=3D"width: 100%;">
        <tbody>
          <tr>

            <td align=3D"right">
              <a href=3D"http://link.xxxpersonals.com/sx/9dfhr4cs2jpxy5sy4j=
p38mmrrk77Ay2ptlsA5n5l9vthzs5tsrb5815grcslgy7xhkmvvzymyb6qkmc65gr017bc7wgdq=
m67qx30rx2knb2d341t24frkc1y7kmmkrsnypz44y4ry49vA0zbfy3pt24t8j56pvbcygjbht8c=
539bq8w00tnb0f5pfcf3n7twmyA92xAxslchqgwf8x78kvkAkA9hys0phg9b29shq3zz51sskgj=
qlffmm7zhrdx7rlkqhs6y18prh587r3mkdr53xx8hfm7fpkvv3r0pfj8zrlwbAgdjxgrw0fj25r=
p484klvw72hvwj4xfx1hm80y44122jlgtdg4mn1fy6b7973fc99wyd0x5kc200d7hw30zA1vl9p=
Ahn3689z0ffm28dy0q1z4f4g7vfxny0s3ktvAj8Asnlr9grfd3c1f6b64jg0dnw75yhsrmjn5vh=
gc30wx33nl2nq99zdzmb0tp6xc0gt2sjh5fcvt55sq508ls84ghmA6r3dsmrlxvq4ykdjc0l4ns=
4A7fykrt6t1g8mwv3znjA5A1hArryjmqr5qb8788tkjznw9mjzfb6sbclyk04ybby8kh2mh34nc=
w5vgc95AybgArcAy00tr6xrgm50hm9ts4jh67sAr5prq1xqtvwvktqt6dsjvj74bgnqsksrAqs1=
81g9b1xwhxzrA84yr7hb28hhmw6v2qxqhmxwlc66rnmt1m81byd72pjs0zf2qq8j6fj6fwv3l28=
xyA5smgygf7c04n79Ay5526vykf66bzAy95k326d9ctwvptps31dcy868cjAg2twnb53rlq3Akf=
3rcn9zs0kfgg7tcfrb1hkb8rh86l9hdh6xnfzdd52Ac7jwk4hfkvrgyh30ylkjr29761ykdA1s2=
vqmvwgf3986v0k7mbAbj11q27h7c0r7Axxd36ppbn4bwdvkj3rwyztjhcsdq5q4gzt8jk1txv1t=
kbjqdjh44y3p0jmlwb41lw9wkgs3zn8j074n33wxv49ydl0dn1hn602qx8dc1bsmgqp790ddkbd=
cfd2pnsbh5p2rx1671fbd63bh3s09djtA82q7dvdms2h4rkxgdxA9w0g3bdbrnlvmrrt7bkcddr=
A9p8p16cyynphqmhpfhl9rbxm1hrcxAjApwpnm2qnrtc3cnxfqtqyjxvs5fAln10tfvt19kzqlq=
xy4k83cxbn8hlA4n2A9btvpxgw62w30f3gwklj9h1t03jAzdkjAkc6ck7cyxk1159z5m2rny34t=
gwbzgq9cp16hhr2qpsl3twfxzqdp66fwp1zmqtrc6mtzqxA6xttq1qlh4frn8d2dp8tkcjn4qvs=
jz7bxvxn222fkzbcn6bqnbmt0tmvk5xAyq9gdwpd01t00ghzlA3n6t5kAqql5krs31hgjn3f1qb=
w8yd74mcmbd1gA8zvh6lm7z4kgxn4lft0hxhzbsq9whApp4mvr4rldzAntjpglhvwyq9bq0lvw0=
5lp34pfgrv0jyjc8bhqryqz9b3zk41q5pp4cAwg5n393/d38"></a>
            </td>
          </tr>
        </tbody>
      </table>
      <span style=3D"color:#888888; font-family: arial; font-size: 12px;">
  XXXPersonals.com &#8226; PO Box 2146 &#8226; Liguanea Post Office &#8226;=
 Kingston 6 &#8226;  Jamaica
      </span>
    </center>
    <img src=3D"http://link.xxxpersonals.com/tg/cbk27nrkklt9ng5h7wskjc0kx71=
4h8hy7vcl1b99k59lp7nc31tgrAp3st1pd1myd05xwjxvt8vj6sAh01m5p12yxvrr5Asrpwpx4m=
wd5y3qs879AfAkjjd7n9t4pf71/b38" height=3D"1" width=3D"1">
   =20
  </body>
</html>
--=_4fe15f765d8e68d30d2f92923456461f--
`
	result := spamcheck.Check(rawEmail, true)

	fmt.Println("Score:", result.Score)

	rules, _ := json.MarshalIndent(result.Rules, "", "  ")
	fmt.Println("Rules:\n", string(rules))

	fmt.Println("Report:\n", result.Report)

	return
}