var codeReader;
var scannerSound = document.getElementById("beepSound");
var selectedDeviceId;
var scanerSources;
var curScanerSourceIdx = 0;

function loadDecoder(videoElementID, isSingleScan, callback, onDevicesLoad) {
    lastVideoElementID = videoElementID;
    lastIsSingleScan = isSingleScan;
    lastCallback = callback;
    lastOnDevicesLoad = onDevicesLoad;

    if(!codeReader){
        codeReader = new ZXing.BrowserBarcodeReader()        
    }else{
        codeReader.reset();
    }

    codeReader.getVideoInputDevices().then((videoInputDevices) => {
            if(videoInputDevices.length > 1){
                $("#iterateScannerBtn").slideDown();
            }

            scanerSources = videoInputDevices;
            selectedDeviceId = videoInputDevices[curScanerSourceIdx].deviceId
            onDevicesLoad(scanerSources)

            if(isSingleScan){
                startSingleScan(videoElementID, callback);
            }else{
                startContinuousScan(videoElementID, callback);
            }

        })
        .catch((err) => {
            console.error(err)
        })
}

function iterateInputDevice(){
    curScanerSourceIdx = (curScanerSourceIdx+1) % scanerSources.length;
    codeReader.reset();
    loadDecoder(lastVideoElementID, lastIsSingleScan, lastCallback, lastOnDevicesLoad); 
}

function resetDecoder(){
    $("#iterateScannerBtn").slideUp();
    codeReader.reset();
}

// TODO: Stop continuous scan

function startContinuousScan(videoElementID, resultCallback){
    codeReader.decodeFromInputVideoDeviceContinuously(selectedDeviceId, videoElementID, (scanResult, err) => {
        if (scanResult) {
            scannerSound.play();
            resultCallback(scanResult.text);
        }

        if (err) {
            // if (err instanceof ZXing.NotFoundException) {console.log('No QR code found.')}
            // if (err instanceof ZXing.ChecksumException) {console.log('A code was found, but it\'s read value was not valid.')}
            // if (err instanceof ZXing.FormatException) {console.log('A code was found, but it was in a invalid format.')}
        }
    })
}

function startSingleScan(videoElementID, resultCallback){
    codeReader.decodeOnceFromVideoDevice(selectedDeviceId, videoElementID).then((result) => {
        // console.log(result);
        scannerSound.play();
        resetDecoder();
        resultCallback(result.text);
    }).catch((err) => {
        // if (err instanceof ZXing.NotFoundException) {console.log('No QR code found.')}
        // if (err instanceof ZXing.ChecksumException) {console.log('A code was found, but it\'s read value was not valid.')}
        // if (err instanceof ZXing.FormatException) {console.log('A code was found, but it was in a invalid format.')}
        // console.error(err);
    })
}
