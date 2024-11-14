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
    "revision": "3fa85c79b475a6086eb112a983353bc1"
  },
  {
    "url": "assets/css/0.styles.b8c02aaa.css",
    "revision": "fdd057c510a0d6562f78205f63b4d380"
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
    "url": "assets/js/10.66f6fd47.js",
    "revision": "9861f1acd35b8cdae58c97c641a020c0"
  },
  {
    "url": "assets/js/11.c1b62882.js",
    "revision": "7d264acc88202f94eeaff5d87f12b22b"
  },
  {
    "url": "assets/js/12.20305a17.js",
    "revision": "f3cfd33f150cd75d30cb5008bed1a848"
  },
  {
    "url": "assets/js/13.685fefd1.js",
    "revision": "329392e254b7c42877bdd9be0419daf3"
  },
  {
    "url": "assets/js/14.e87425e5.js",
    "revision": "7e6e78346124412195991dd1ef11bfae"
  },
  {
    "url": "assets/js/15.ebf38a50.js",
    "revision": "21e30b797528e6510270412867bde0ba"
  },
  {
    "url": "assets/js/16.eb9cfdc4.js",
    "revision": "6fa626a10aef3866ac6390b31d5bd998"
  },
  {
    "url": "assets/js/17.3fb5cdee.js",
    "revision": "666592bb25546b198f987ac0b84a4d6a"
  },
  {
    "url": "assets/js/18.08e99a96.js",
    "revision": "7d111daf9133d4f48a6b903d09763fa1"
  },
  {
    "url": "assets/js/19.19f206bb.js",
    "revision": "61a0efe9ad2a0b831e70315ea4cae3ba"
  },
  {
    "url": "assets/js/2.a17e1230.js",
    "revision": "daa6b6ed292696bf0a2e3aa0e1b1153a"
  },
  {
    "url": "assets/js/20.c005b3ae.js",
    "revision": "426cc3dd3817b2835bc049c686fad700"
  },
  {
    "url": "assets/js/21.9aa28d1d.js",
    "revision": "e06926d5339e07721f478f686c9a0507"
  },
  {
    "url": "assets/js/22.780c2ff7.js",
    "revision": "69681d86119ec3370c19607fea6520ed"
  },
  {
    "url": "assets/js/23.74853ab5.js",
    "revision": "84ea8034415030012ee9aababfc55c86"
  },
  {
    "url": "assets/js/24.0a0bcb8b.js",
    "revision": "0a29e868ac84e7d34723ef9e9be2efca"
  },
  {
    "url": "assets/js/26.a388051f.js",
    "revision": "9052f0ce80ac768f33661d8b7dc67bcd"
  },
  {
    "url": "assets/js/3.37ff259d.js",
    "revision": "9d2b51430537848cac2ea5124482a938"
  },
  {
    "url": "assets/js/4.62df5e3a.js",
    "revision": "668bc91fb4500762c33bec844d358f72"
  },
  {
    "url": "assets/js/5.01993805.js",
    "revision": "6d5176ba3b4c2fa125c6a9aeb57b6aa7"
  },
  {
    "url": "assets/js/6.d2d7c328.js",
    "revision": "52d18539e934aaf4d345354d67942f01"
  },
  {
    "url": "assets/js/7.c09e37a9.js",
    "revision": "bdfef6f12cfe483be868002afeba6c65"
  },
  {
    "url": "assets/js/8.57abb3f8.js",
    "revision": "a323d6e61a5078e66e9a97db0719672c"
  },
  {
    "url": "assets/js/9.a50bcd73.js",
    "revision": "8d744c4bc62ac7fd91f68755a0f6239d"
  },
  {
    "url": "assets/js/app.337080e8.js",
    "revision": "ed294effcef80041f638b26acd85912f"
  },
  {
    "url": "conclusion/index.html",
    "revision": "e2896f94e5e176ba100a21d2ed174182"
  },
  {
    "url": "design/index.html",
    "revision": "432a8fa2e5eb902477ed6b11a1be4276"
  },
  {
    "url": "index.html",
    "revision": "1d9e62d51ef96a13fc3c3928eda6836b"
  },
  {
    "url": "intro/index.html",
    "revision": "ce59ebe51b998ad0bb151a230ab0a876"
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
    "revision": "bd6be0336d3a567e79766742cafc235f"
  },
  {
    "url": "requirements/stakeholders-needs.html",
    "revision": "fe1c8f097d4d13a085ad767c61d3f4eb"
  },
  {
    "url": "requirements/state-of-the-art.html",
    "revision": "2050b41c412436208dbbf8625dd29b06"
  },
  {
    "url": "software/index.html",
    "revision": "fbb89d3bd48d7aac4080d3a0778c7b4b"
  },
  {
    "url": "test/index.html",
    "revision": "2e35338c8d6dfb76cc8aa7305813da3c"
  },
  {
    "url": "use cases/index.html",
    "revision": "f452106989386ef62270eb4544d5e036"
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
