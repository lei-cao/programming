function loadEbiten() {
    var doc = document.getElementById('ebitenIframe').contentWindow.document
    doc.open()
    doc.write('<html><head><title></title></head><body><script src="/scripts/algoman.js?" + Date.now()></script></body></html>')
    doc.close()
}