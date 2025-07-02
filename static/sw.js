const fileMap = new Map();

self.addEventListener('install', () => self.skipWaiting());
self.addEventListener('activate', () => self.clients.claim());

self.addEventListener('message', e => {
    if (e.data?.type === 'FILES') {
        for (const [path, buffer] of Object.entries(e.data.files)) {
            fileMap.set('/virtual/' + path, buffer);
        }
        // ✅ Acknowledge receipt
        e.source.postMessage({ type: 'FILES_ACK' });
    }
});

self.addEventListener('fetch', e => {
    const path = new URL(e.request.url).pathname;
    if (fileMap.has(path)) {
        e.respondWith(new Response(fileMap.get(path), {
            headers: { 'Content-Type': getContentType(path) }
        }));
    }
});

function getContentType(path) {
    if (path.endsWith('.html')) return 'text/html';
    if (path.endsWith('.js')) return 'application/javascript';
    if (path.endsWith('.css')) return 'text/css';
    if (path.endsWith('.png')) return 'image/png';
    if (path.endsWith('.jpg') || path.endsWith('.jpeg')) return 'image/jpeg';
    if (path.endsWith('.gif')) return 'image/gif';
    if (path.endsWith('.mp3')) return 'audio/mpeg';
    return 'application/octet-stream';
}