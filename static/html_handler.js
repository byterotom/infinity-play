import { unzipSync } from 'https://cdn.skypack.dev/fflate';

await navigator.serviceWorker.register('/static/sw.js', { scope: '/' });
await navigator.serviceWorker.ready;

if (!navigator.serviceWorker.controller) {
    await new Promise(resolve =>
        navigator.serviceWorker.addEventListener('controllerchange', resolve, { once: true })
    );
}

const res = await fetch(`/game/html/${gameId}`);
const zipBuffer = await res.arrayBuffer();
const zip = unzipSync(new Uint8Array(zipBuffer));

const fileMap = {};
for (const [name, data] of Object.entries(zip)) {
    fileMap[name] = data.buffer;
}

const swReady = new Promise(resolve => {
    navigator.serviceWorker.addEventListener('message', event => {
        if (event.data?.type === 'FILES_ACK') resolve();
    }, { once: true });
});

navigator.serviceWorker.controller.postMessage({ type: 'FILES', files: fileMap });
await swReady;

const entryFile = Object.keys(fileMap).find(name => name.endsWith('index.html'));
document.getElementById('gameObject').src = '/virtual/' + entryFile;
