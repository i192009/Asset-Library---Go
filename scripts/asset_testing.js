const fs = require('fs');
const path = require('path');

const addAsset = async (
    filename, filesize, fileType, thumbnail, thumbnailType,
    thumbnailSize, classId, openId, appId, authToken
) => {
    const url = 'https://dev.zixel.cn/api/assetManage/v1/backend/asset/add';
    const bodyData = {
        "filename": filename,
        "filesize": filesize,
        "type": fileType,
        "thumbnail": thumbnail,
        "thumbnailType": thumbnailType,
        "thumbnailSize": thumbnailSize,
        "class": classId,
        "platform": "1",
        "permissionType": "1",
        "title": "少世斗手四",
        "description": "工构民由产加体眼不按团格酸极。。",
        "owner": openId,
        "external": {}
    };

    try {
        const response = await fetch(url, {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json',
                'Zixel-Open-Id': openId,
                'Zixel-Application-Id': appId,
                'Zixel-Auth-Token': authToken,
            },
            body: JSON.stringify(bodyData)
        });

        if (!response.ok) {
            throw new Error('Network response was not ok');
        }

        const data = await response.json();
        console.log('Success:', data);

        return {
            assetId: data.assetId,
            assetUploadUrl: data.assetUploadUrl,
            thumbUploadUrl: data.thumbUploadUrl
        };
    } catch (error) {
        console.error('Error:', error);
        throw error;
    }
};

const getAsset = async (assetId) => {
    const url = `https://dev.zixel.cn/api/assetManage/v1/backend/asset/get/1/${assetId}`;

    try {
        const response = await fetch(url, {
            method: 'GET',
            headers: {
                'Content-Type': 'application/json',
                'Zixel-Open-Id' : openId,
                'Zixel-Application-Id' : appId,
                'Zixel-Auth-Token' : authToken,
            }
        });

        if (!response.ok) {
            throw new Error('Network response was not ok');
        }

        const data = await response.json();
        console.log('Success:', data);
        return data;
    } catch (error) {
        console.error('Error:', error);
        throw error;
    }
};

const updateAssetStatus = async (assetId) => {
    const url = 'https://dev.zixel.cn/api/assetManage/v1/backend/asset/assetUploaded/1/' + assetId;

    try {
        const response = await fetch(url, {
            method: 'POST',
            headers: {
                'Asset-Id': assetId,
                'Zixel-Open-Id' : openId,
                'Zixel-Application-Id' : appId,
                'Zixel-Auth-Token' : authToken,
            },
        });

        if (!response.ok) {
            throw new Error('Network response was not ok');
        }

        console.log('Success:');
        return "Success";
    } catch (error) {
        console.error('Error:', error);
        throw error;
    }
};

const updateAssetThumbnail = async (assetId, thumbnailSize, thumbnailType) => {
    const url = `https://dev.zixel.cn/api/assetManage/v1/backend/asset/update/1/${assetId}/thumbnail`;
    const bodyData = {
        "thumbnailSize": thumbnailSize,
        "thumbnailType": thumbnailType
    };

    try {
        const response = await fetch(url, {
            method: 'PUT',
            headers: {
                'Content-Type': 'application/json',
                'Zixel-Open-Id' : openId,
                'Zixel-Application-Id' : appId,
                'Zixel-Auth-Token' : authToken,
            },
            body: JSON.stringify(bodyData)
        });

        if (!response.ok) {
            throw new Error('Network response was not ok');
        }

        const data = await response.json();
        console.log('Success:', data);

        return data.thumbnail;  // Assuming that response contains a 'thumbnail' property
    } catch (error) {
        console.error('Error:', error);
        throw error;
    }
};

const readFile = (dir, fileName) => {
    return new Promise((resolve, reject) => {
        const filePath = path.join(dir, fileName);

        fs.readFile(filePath, (err, data) => {
            if (err) {
                console.error('Failed to read file:', err);
                return reject(err);
            }

            fs.stat(filePath, (err, stats) => {
                if (err) {
                    console.error('Failed to retrieve file stats:', err);
                    return reject(err);
                }
                console.log(`File size: ${stats.size}`);
                console.log(`File name: ${fileName}`);
                resolve({
                    filename: fileName,
                    filesize: stats.size,
                    file: data
                });
            });
        });
    });
};

const sendFileToUrl = async (url, file, contentType) => {
    try {

        // Create HTTP request and send file data
        const response = await fetch(url, {
            method: 'PUT',
            body: file,
            headers: { 'Content-Type': contentType }
        });

        if (!response.ok) {
            console.log('Response status:', response.status);
            console.log('Response status text:', response.statusText);
            throw new Error('Network response was not ok: ' + response.statusText);
        }
        console.log('Success');
        return "Success";
    } catch (error) {
        console.error('Error:', error);
        throw error;
    }
};

const manageAssetUpload = async (
    dir, fileName, thumbnailDir, thumbnailFileName,
    fileType, thumbnailType,
    classId, openId, appId, authToken
) => {
    try {
        // 1. Read file
        const { filename, filesize, file } = await readFile(dir, fileName);
        const { filename: thumbnailFilename, filesize: thumbnailFilesize, file: thumbnailFile } = await readFile(thumbnailDir, thumbnailFileName);

        // 2. Add asset
        const { assetId, assetUploadUrl, thumbUploadUrl } = await addAsset(
            filename, filesize, fileType,
            thumbnailFilename, thumbnailType, thumbnailFilesize,
            classId, openId, appId, authToken
        );

        console.log(`Asset ID: ${assetId}`);

        // 3. Upload files
        console.log('Upload to url: ' + assetUploadUrl)
        await sendFileToUrl(assetUploadUrl, file, fileType);
        console.log('File uploaded successfully');

        console.log('Upload to url: ' + thumbUploadUrl)
        await sendFileToUrl(thumbUploadUrl, thumbnailFile, thumbnailType);
        console.log('Thumbnail uploaded successfully');

        await updateAssetStatus(assetId);
        console.log('Asset status updated successfully');

        await getAsset(assetId);
        console.log('Asset retrieved successfully');

        const thumbnailUrl = await updateAssetThumbnail(assetId, filesize);
        console.log('Thumbnail Url retrieved successfully!');

        console.log('Upload to url: ' + thumbnailUrl)
        await sendFileToUrl(thumbnailUrl, file, "");
        console.log('Thumbnail uploaded successfully');

        await getAsset(assetId);
                console.log('Asset retrieved successfully');

    } catch (error) {
        console.error('Error in manageAssetUpload:', error);
    }
};

const manageUpdateThumbnail = async (assetId, dir, filename, fileType, openId, appId, authToken) => {
    try {
        // 1. Read file
        const { filename: readFilename, filesize, file } = await readFile(dir, filename);

        // 2. Update thumbnail
        const thumbnail = await updateAssetThumbnail(assetId, filesize, fileType);

        // 3. Send file to URL
        const uploadedFileResponse = await sendFileToUrl(thumbnail, file, fileType);

        console.log('File uploaded successfully:', uploadedFileResponse);
    } catch (error) {
        console.error('Error in manageUpdateThumbnail:', error);
    }
};

const openId = "oi_45115871488";
const appId = "collaborate";
const authToken = "ut_UZ75HO59V0VBC300T5ZskJrX44eCWq5I";

const assetId = "652640dcaaf845457bdb56b4"
//const classId = "6513ec2fdb86a71e50a9667d"; // Image
const classId = "6513ec2fdb86a71e50a9667f"; // Audio
const dir = './testAssets';
const fileName = 'file_example_MP3_700KB.mp3';
const fileType = 'audio/mp3';
const thumbnailDir = './testAssets';
const thumbnailFileName = 'sample.svg';
const thumbnailType = 'image/svg+xml';

manageAssetUpload(dir, fileName, thumbnailDir, thumbnailFileName, fileType, thumbnailType, classId, openId, appId, authToken);
manageUpdateThumbnail(assetId, thumbnailDir, thumbnailFileName, thumbnailType, openId, appId, authToken);