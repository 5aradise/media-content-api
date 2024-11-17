/**
 * Welcome to your Workbox-powered service worker!
 *
 * You'll need to register this file in your web app and you should
 * disable HTTP caching for this file too.
 * See https://goo.gl/nhQhGp
 *
 * The rest of the code is auto-generated. Please don't update this file
 * directly; instead, make changes to your Workbox build configuration
 * and re-run your build process.
 * See https://goo.gl/2aRDsh
 */

importScripts("https://storage.googleapis.com/workbox-cdn/releases/4.3.1/workbox-sw.js");

self.addEventListener('message', (event) => {
  if (event.data && event.data.type === 'SKIP_WAITING') {
    self.skipWaiting();
  }
});

/**
 * The workboxSW.precacheAndRoute() method efficiently caches and responds to
 * requests for URLs in the manifest.
 * See https://goo.gl/S9QRab
 */
self.__precacheManifest = [
  {
    "url": "404.html",
    "revision": "46cfdfb9df9ef2f5661e502d8f5636dc"
  },
  {
    "url": "assets/css/0.styles.b8c02aaa.css",
    "revision": "fdd057c510a0d6562f78205f63b4d380"
  },
  {
    "url": "assets/img/media_content-delete.172af4a7.png",
    "revision": "172af4a72a1345eed58e0307fea426c8"
  },
  {
    "url": "assets/img/media_content-delete.d49ba168.png",
    "revision": "d49ba16825f53968a8cabc2305d67bd7"
  },
  {
    "url": "assets/img/media_content-empty-required-field.b3ce7ee5.png",
    "revision": "b3ce7ee555eebc567125a8b6e26d2e7f"
  },
  {
    "url": "assets/img/media_content-get-id.9dc94daf.png",
    "revision": "9dc94dafab7bed6eb557fdc68dfbec2a"
  },
  {
    "url": "assets/img/media_content-get-id.b4ead904.png",
    "revision": "b4ead904b86b99b78844b5c215299fcc"
  },
  {
    "url": "assets/img/media_content-get-user_id.03aa310e.png",
    "revision": "03aa310e5a912c737c6d33e26f6f0632"
  },
  {
    "url": "assets/img/media_content-get.60a642e8.png",
    "revision": "60a642e8722361ce22fd4cfa2bbe5c28"
  },
  {
    "url": "assets/img/media_content-get.96d40cd8.png",
    "revision": "96d40cd893529ff06fd8b032638f8500"
  },
  {
    "url": "assets/img/media_content-invalid-media-content-type.3f7d3566.png",
    "revision": "3f7d3566feee55ff8095649bed511c00"
  },
  {
    "url": "assets/img/media_content-post.62040d20.png",
    "revision": "62040d200140ae3100357bf8d4edf7d9"
  },
  {
    "url": "assets/img/media_content-post.ee42a1d3.png",
    "revision": "ee42a1d30b52ed3235fca4ab38e9d95f"
  },
  {
    "url": "assets/img/media_content-unfound-by-id.0471ccef.png",
    "revision": "0471ccef53448efb2d2c125a6e44173f"
  },
  {
    "url": "assets/img/media_content-user-id-violation.22490785.png",
    "revision": "224907850c052d0481cc540f54e2bc02"
  },
  {
    "url": "assets/img/relational_schema.ce95178e.png",
    "revision": "ce95178e03036cd4d2d2ffe1a92aaba3"
  },
  {
    "url": "assets/img/search.83621669.svg",
    "revision": "83621669651b9a3d4bf64d1a670ad856"
  },
  {
    "url": "assets/img/swagger.356b9e77.png",
    "revision": "356b9e77d45c146bdc0b5490c585e885"
  },
  {
    "url": "assets/img/users-delete.19ae3b95.png",
    "revision": "19ae3b95ee8645100dd68e66f30a4461"
  },
  {
    "url": "assets/img/users-delete.9a714eb8.png",
    "revision": "9a714eb81aeed2e086ad588292a1aef1"
  },
  {
    "url": "assets/img/users-email-unique-violation-post.50cd02d3.png",
    "revision": "50cd02d370f2c02a509f9aa9132019e9"
  },
  {
    "url": "assets/img/users-email-unique-violation-put.fec67f82.png",
    "revision": "fec67f821b543656e76ff1f54d4c20fe"
  },
  {
    "url": "assets/img/users-empty-required-field.d310ff9b.png",
    "revision": "d310ff9b95119498b093609ad2baf7ac"
  },
  {
    "url": "assets/img/users-get-id.710846a7.png",
    "revision": "710846a7b778b17eb8e2a57b57517ca3"
  },
  {
    "url": "assets/img/users-get-id.76402a6b.png",
    "revision": "76402a6bba94b4286d15edac3154422b"
  },
  {
    "url": "assets/img/users-get.01a914e1.png",
    "revision": "01a914e1a23c4f07a5c90fbc4c3e2c50"
  },
  {
    "url": "assets/img/users-get.0b7cb200.png",
    "revision": "0b7cb2000eac6f6327257f3e6b5b03bb"
  },
  {
    "url": "assets/img/users-invalid-email-post.9f66f092.png",
    "revision": "9f66f09238d99b60c3052e904896b46d"
  },
  {
    "url": "assets/img/users-invalid-email-put.4663cc03.png",
    "revision": "4663cc0337ba304771e9ac3dca9507e3"
  },
  {
    "url": "assets/img/users-post.ca640af2.png",
    "revision": "ca640af2f6543f02e2e6ff00a1530916"
  },
  {
    "url": "assets/img/users-post.ddcc8efa.png",
    "revision": "ddcc8efa4e99426b24f671db856e6b8a"
  },
  {
    "url": "assets/img/users-put.14b1db71.png",
    "revision": "14b1db71a5225921461c3481ce45e226"
  },
  {
    "url": "assets/img/users-put.69642460.png",
    "revision": "69642460ed5d1e7bd0d1bf58c5f37718"
  },
  {
    "url": "assets/img/users-unfound-by-id-get.5cdef173.png",
    "revision": "5cdef173682b73aaf8fba6f0f321ad88"
  },
  {
    "url": "assets/img/users-unfound-by-id-put.3cdb03dc.png",
    "revision": "3cdb03dceec5287e3d51849bb64486e6"
  },
  {
    "url": "assets/js/10.ecf0dbb3.js",
    "revision": "85bdde4a8145d71de375f4779db1e6f5"
  },
  {
    "url": "assets/js/11.5338dd9c.js",
    "revision": "a8de80e247bd8b7653e1550e928c175f"
  },
  {
    "url": "assets/js/12.6e0f26d3.js",
    "revision": "69c3baa453f21a4c94d39d9a6f1f17f0"
  },
  {
    "url": "assets/js/13.467d8ba9.js",
    "revision": "a25dd56e7eb01a417277f08e2c5bff6a"
  },
  {
    "url": "assets/js/14.83187df0.js",
    "revision": "7f6e601bbdd3b67369bca34394ff6fe7"
  },
  {
    "url": "assets/js/15.0c0ca6f1.js",
    "revision": "bb402a5327ecf829a03896c9dc8fdfd3"
  },
  {
    "url": "assets/js/16.9b09b718.js",
    "revision": "fd3574b6d65ff227a5df885e0a89c4ee"
  },
  {
    "url": "assets/js/17.8400d547.js",
    "revision": "2b42d45afa790f0ec4973f9e53381757"
  },
  {
    "url": "assets/js/18.fdd2db53.js",
    "revision": "5a7a576f0b077a8d4e0566a98a1da373"
  },
  {
    "url": "assets/js/19.7b1b496e.js",
    "revision": "602cf302fc47993cfca56ae8e24ef321"
  },
  {
    "url": "assets/js/2.f2c4eaec.js",
    "revision": "46391393d22bb2ea32e50361654bdd26"
  },
  {
    "url": "assets/js/20.feab4e5b.js",
    "revision": "01cff5922d2ed707db7f9cfd40ee17bf"
  },
  {
    "url": "assets/js/21.afdf077d.js",
    "revision": "9d36239886ca9dcb7a34f4562ff09619"
  },
  {
    "url": "assets/js/22.0d6c67ba.js",
    "revision": "a50764908759f2a1765b14660441f5d1"
  },
  {
    "url": "assets/js/23.c5bdcb79.js",
    "revision": "6ca2612322c960b68f571f60bc3feac9"
  },
  {
    "url": "assets/js/24.c58aa348.js",
    "revision": "12fb04e3ce83c6856f9e9ab8fd04ae62"
  },
  {
    "url": "assets/js/26.6f18bd18.js",
    "revision": "30408fffa4c071e4a0a6dc1429cb5203"
  },
  {
    "url": "assets/js/3.7be694a3.js",
    "revision": "147695dfad44720cf9f5cbdfe089bc2d"
  },
  {
    "url": "assets/js/4.92b89156.js",
    "revision": "563dbc2f9f520e4986de7e6b7ae2e4ca"
  },
  {
    "url": "assets/js/5.f1bd00d8.js",
    "revision": "05587c9e4f79fd6c443b8988e9f6742d"
  },
  {
    "url": "assets/js/6.fcf25c65.js",
    "revision": "bd98efe375040c22964d66fda5179ae0"
  },
  {
    "url": "assets/js/7.1b47a0a0.js",
    "revision": "eda41cab6436e70b95b5f2929f04bcc3"
  },
  {
    "url": "assets/js/8.0cb558aa.js",
    "revision": "2f84ce7cd09841124d56ef663d8f3b23"
  },
  {
    "url": "assets/js/9.9cb5584e.js",
    "revision": "919230b90a4961b9a8326724d459c688"
  },
  {
    "url": "assets/js/app.b3233e4e.js",
    "revision": "89f7f18ede015d4eed95952a032e24ff"
  },
  {
    "url": "conclusion/index.html",
    "revision": "e335b87b530aec68de3d124604e868ec"
  },
  {
    "url": "design/index.html",
    "revision": "5ca4294dd7310330d0e91574246d7dfe"
  },
  {
    "url": "index.html",
    "revision": "1320a6ddfdae3f0f291e3baf5eb2748f"
  },
  {
    "url": "intro/index.html",
    "revision": "85297d0aa380dd3b0e7a33be1d53a064"
  },
  {
    "url": "license.html",
    "revision": "c6ae94487788630ac26fdec72cecb7f2"
  },
  {
    "url": "myAvatar.png",
    "revision": "b76db1e62eb8e7fca02a487eb3eac10c"
  },
  {
    "url": "requirements/index.html",
    "revision": "357b45d4ad9077cc2d7de566593998fd"
  },
  {
    "url": "requirements/stakeholders-needs.html",
    "revision": "f6834a425fdd1452943d87ad27bc56a0"
  },
  {
    "url": "requirements/state-of-the-art.html",
    "revision": "fb596f377e34f5c5bc36c2dead704b67"
  },
  {
    "url": "software/index.html",
    "revision": "a889f2dc3d3b2465dc62efa8e9706213"
  },
  {
    "url": "test/index.html",
    "revision": "a572c5363ae60701348429e2fdae970e"
  },
  {
    "url": "use cases/index.html",
    "revision": "7e874214bf69938070ccf53d831d0197"
  }
].concat(self.__precacheManifest || []);
workbox.precaching.precacheAndRoute(self.__precacheManifest, {});
addEventListener('message', event => {
  const replyPort = event.ports[0]
  const message = event.data
  if (replyPort && message && message.type === 'skip-waiting') {
    event.waitUntil(
      self.skipWaiting().then(
        () => replyPort.postMessage({ error: null }),
        error => replyPort.postMessage({ error })
      )
    )
  }
})
