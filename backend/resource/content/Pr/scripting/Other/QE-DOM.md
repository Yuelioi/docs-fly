---
title: QE DOM
order: 2
---

---
title: QE DOM
order: 5
category:

- PR脚本

---

QE Dom 是指加载额外的函数与属性（你就这么理解吧），需要先打开 QE 开关

https://www.youtube.com/watch?v=Q7UTnEiUZ5o

https://www.youtube.com/watch?v=Q7UTnEiUZ5o

## QE

var qe = app.enableQE();

## global

qe.exit();

qe.getSequencePresets();

qe.language;

qe.location;

qe.name;

qe.newProject();

qe.open();

qe.outputToConsole();

qe.platform;

qe.startPlayback();

qe.stop();

qe.stopPlayback();

## project

var qeProj = qe.project;

qeProj.close();

qeProj.currentRendererName;

qeProj.deletePreviewFiles();

qeProj.flushCache();

qeProj.getAudioEffectByName();

qeProj.getAudioEffectList();

qeProj.getAudioTransitionByName();

qeProj.getAudioTransitionList();

qeProj.getBinAt();

qeProj.getItemAt();

qeProj.getSequenceAt();

qeProj.getSequenceItemAt();

qeProj.getVideoEffectByName();

qeProj.getVideoEffectList();

qeProj.getVideoTransitionByName();

qeProj.getVideoTransitionList();

qeProj.import();

qeProj.importAEComps();

qeProj.importAllAEComps();

qeProj.importFiles();

qeProj.importProject();

qeProj.name;

qeProj.newBarsAndTone();

qeProj.newBin();

qeProj.newBlackVideo();

qeProj.newCaption();

qeProj.newSequence();

qeProj.newTransparentVideo();

qeProj.numBins;

qeProj.numItems;

qeProj.numSequenceItems;

qeProj.numSequences;

qeProj.path;

qeProj.redo();

qeProj.undo();

## source

var qeSource = qe.source;

qeSource.clearInPoint();

qeSource.clearOutPoint();

qeSource.duration;

qeSource.filePath;

qeSource.name;

qeSource.setAudioInPoint();

qeSource.setAudioOutPoint();

qeSource.setDuration();

qeSource.setInPoint();

qeSource.setOutPoint();

qeSource.setVideoInPoint();

qeSource.setVideoOutPoint();

## PROJECT ITEMS

// properties

QEProjectItem.clip;

QEProjectItem.filePath;

QEProjectItem.name;

// methods

QEProjectItem.automateToSequence(object, number, number, number, number);

QEProjectItem.containsSpeechTrack();

QEProjectItem.createProxy(string, string);

QEProjectItem.getMetadataSize();

QEProjectItem.isAudioConforming();

QEProjectItem.isAudioPeakGenerating();

QEProjectItem.isIndexing();

QEProjectItem.isOffline();

QEProjectItem.isPending();

QEProjectItem.linkMedia(string, bool);

QEProjectItem.openInSource();

QEProjectItem.rename(assetName);

QEProjectItem.setOffline();

## SEQUENCES

// properties

QESequence.CTI;

QESequence.audioDisplayFormat;

QESequence.audioFrameRate;

QESequence.editingMode;

QESequence.fieldType;

QESequence.guid;

QESequence.inPoint;

QESequence.multicam;

QESequence.name;

QESequence.numAudioTracks;

QESequence.numVideoTracks;

QESequence.outPoint;

QESequence.par;

QESequence.player;

QESequence.presetList;

QESequence.previewPresetCodec;

QESequence.previewPresetPath;

QESequence.useMaxBitDepth;

QESequence.useMaxRenderQuality;

QESequence.videoDisplayFormat;

QESequence.videoFrameRate;

QESequence.workInPoint;

QESequence.workOutPoint;

// methods

QESequence.addTracks(number, number, number, ...);

QESequence.close();

QESequence.createExportJob(string, string);

QESequence.deletePreviewFiles(string, string);

QESequence.exportDirect(string, string, bool);

QESequence.exportFrameBMP(string, string, string);

QESequence.exportFrameDPX(string, string, string);

QESequence.exportFrameGIF(string, string, string);

QESequence.exportFrameJPEG(string, string, string);

QESequence.exportFramePNG(string, string, string);

QESequence.exportFrameTIFF(string, string, string);

QESequence.exportFrameTarga(string, string, string);

QESequence.exportToAME(string, string, bool);

QESequence.extract(string, string);

QESequence.flushCache();

QESequence.getAudioTrackAt(number);

QESequence.getEmptyBarTimes();

QESequence.getExportComplete();

QESequence.getExportFileExtension(string);

QESequence.getGreenBarTimes();

QESequence.getRedBarTimes();

QESequence.getVideoTrackAt(number);

QESequence.getYellowBarTimes();

QESequence.isIncompleteBackgroundVideoEffects();

QESequence.isOpen();

QESequence.left(string, string);

QESequence.lockTracks(string, bool);

QESequence.makeCurrent();

QESequence.muteTracks(string, bool);

QESequence.razor(string, bool, bool);

QESequence.removeAudioTrack(number);

QESequence.removeEmptyAudioTracks();

QESequence.removeEmptyVideoTracks();

QESequence.removeTracks(number, number, number, ...);

QESequence.removeVideoTrack(number);

QESequence.renderAll();

QESequence.renderAudio();

QESequence.renderPreview();

QESequence.setAudioDisplayFormat(number);

QESequence.setAudioFrameRate(number);

QESequence.setCTI(string);

QESequence.setInOutPoints(string, string, bool);

QESequence.setInPoint(string, bool, bool);

QESequence.setOutPoint(string, bool, bool);

QESequence.setPreviewFrameSize(number, number);

QESequence.setPreviewPresetPath(string);

QESequence.setUseMaxBitDepth(bool);

QESequence.setUseMaxRenderQuality(bool);

QESequence.setVideoDisplayFormat(number);

QESequence.setVideoFrameSize(number, number);

QESequence.setWorkInOutPoints(string, string, bool);

QESequence.setWorkInPoint(string, bool);

QESequence.setWorkOutPoint(string, bool);

QESequence.syncLockTracks(string, bool);

## TRACKS

// properties

QETrack.id;

QETrack.index;

QETrack.name;

QETrack.numComponents;

QETrack.numItems;

QETrack.numTransitions;

QETrack.type;

// methods

QETrack.addAudioEffect(object, number, bool);

QETrack.getComponentAt(number);

QETrack.getItemAt(number);

QETrack.getTransitionAt(number);

QETrack.insert(object, string, bool, bool, bool, bool);

QETrack.isLocked();

QETrack.isMuted();

QETrack.isSyncLocked();

QETrack.overwrite(object, string, bool, bool, bool);

QETrack.razor(string, bool, bool);

QETrack.setLock(bool);

QETrack.setMute(bool);

QETrack.setName(string);

QETrack.setSyncLock(bool);
