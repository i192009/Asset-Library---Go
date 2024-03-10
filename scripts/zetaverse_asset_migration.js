const { MongoClient } = require('mongodb');

// MongoDB connection string
const uri = "mongodb://127.0.0.1:27017/";

// Create a new MongoClient
const client = new MongoClient(uri, { useNewUrlParser: true, useUnifiedTopology: true });

function getClassName(type, subType) {
  // Logic to determine class name based on type and subType
  switch(type) {
    case 1: return "image";
    case 2: return "video";
    case 3: return 'audio';
    case 6: return "model";
    case 7: {
      switch(subType) {
        case 1: return "scene";
        case 2: return "interactive";
        case 3: return "core";
        default: {
            throw new Error(`Unsupported type: ${subType}`)
        }
      }
    }
    default: throw new Error(`Unsupported type: ${type}`);
  }
}

async function getClassId(appId, className, classCollection) {
  // Fetch the classId from the 'class' collection based on appId and className
  const classDoc = await classCollection.findOne({ appId, name: className });
  if (!classDoc) throw new Error(`Class not found for appId:${appId}, name:${className}`);
  return classDoc._id.toString();
}

function getType(type, subtype, filename) {
    let fileFormat;
    if (filename) {
        fileFormat = filename.split('.').pop();
    }

    switch (type) {
        case 1:
            switch (subtype) {
                case 1: return "image/jpeg";
                case 2: return "image/png";
                case 3: return "image/svg+xml";
                case 4: return "image/gif";
                default: throw new Error(`Unsupported subtype: ${subtype}`);
            }
        case 2:
            switch (subtype) {
                case 1: return "audio/mp3";
                case 2: return "audio/mp4";
                case 3: return "audio/ogg";
                default: throw new Error(`Unsupported subtype: ${subtype}`);
            }
        case 3:
            switch (subtype) {
                case 1: return "video/avi";
                case 2: return "video/mp4";
                default: throw new Error(`Unsupported subtype: ${subtype}`);
            }
        case 6:
            switch (subtype) {
                case 1: return "application/zip";
                case 2: return "application/x-tar";
                case 3:
                case 4:
                case 5:
                case 6: return "application/octet-stream";
                default: throw new Error(`Unsupported subtype: ${subtype}`);
            }
        case 7:
            if (fileFormat === "zip") {
                return "application/zip";
            } else if (fileFormat === "tar") {
                return "application/x-tar";
            } else {
                throw new Error(`Unsupported file format: ${fileFormat}`);
            }
        default: throw new Error(`Unsupported type: ${type}`);
    }
}

async function getTagIds(tags, appId, tagCollection) {
  const tagIds = [];
  if (tags == null){
    return tagIds;
  }
  for (let tagName of tags) {
    const tagDoc = await tagCollection.findOne({ name: tagName.trim(), appId });
    if (!tagDoc) throw new Error(`Tag not found for appId:${appId}, name:${tagName}`);
    tagIds.push(tagDoc._id.toString());
  }
  return tagIds;
}

async function transformAndInsertData() {
  try {
    // Connect to the MongoDB cluster
    await client.connect();

    // Select the 'zetaverse2' database and 'assets' collection
    const sourceDb = client.db('zetaverse2');
    const sourceCollection = sourceDb.collection('assets');

    // Select the 'assetLibrary' database and 'publicAsset' collection
    const destDb = client.db('assetLibrary');
    const destCollection = destDb.collection('publicAsset');
    const classCollection = destDb.collection('class');
    const tagCollection = destDb.collection('tag');

    // Get the documents from source collection where status is 3
    const cursor = sourceCollection.find({ status: 3 });

    // Iterate over all documents from source collection
    while (await cursor.hasNext()) {
      const doc = await cursor.next();
      console.log("Migrating : " + doc._id)

      const className = getClassName(doc.type, doc.subtype);
      const classId = await getClassId("zetaverse", className, classCollection);
      const typeValue = getType(doc.type, doc.subtype, doc.filename);
      const tagIds = await getTagIds(doc.tags, "zetaverse", tagCollection); // Assume appId is 'assetLibrary'


      let external = {}; // default value

      // Check if 'external' field exists in the source document
      if ('external' in doc) {
          external = doc.external;
      }

      // Transform data according to the desired format
      const transformedDoc = {
        _id: doc._id,
        filename: doc.filename,
        filesize: doc.filesize,
        type: typeValue, // Need to modify according to the type mapping logic
        class: classId, // Your class value
        tags: tagIds, // Transform tags to your new format
        title: doc.title,
        instanceId: "", // Your instanceId value
        appId: "zetaverse", // Your appId value
        description: "", // Your description value
        thumbnail: "dev_2/"+ doc.thumbnail,
        url: "dev_2/" + doc.url, // This may need a transformation if not a full URL
        permissionType: "",
        status: doc.status,
        source: doc.source,
        external: external,
        creator: doc.user, // Your creator value
        owner: doc.user,
        createTime: doc.createTime,
        updateTime: doc.updateTime,
      };

      let destCollectionName;
      switch (doc.permissionType) {
        case "1":
        case "共有":
            transformedDoc.permissionType = "1";
            destCollectionName = 'publicAsset';
            break;
        case "2":
        case "企业":
            transformedDoc.permissionType = "2";
            destCollectionName = 'tenantAssets';
            break;
        case "3":
        case "私有":
            transformedDoc.permissionType = "3";
            destCollectionName = 'privateAssets';
            break;
        default:
            transformedDoc.permissionType = "1";
            destCollectionName = 'publicAsset';
      }

        const destCollection = destDb.collection(destCollectionName);
        const filter = { _id: transformedDoc._id }; // This assumes _id is what determines uniqueness
        const update = { $set: transformedDoc };

        await destCollection.updateOne(filter, update, { upsert: true });
    }

  } finally {
    // Close the connection to the MongoDB cluster
    await client.close();
  }
}

// Execute the function to transform and insert data
transformAndInsertData().catch(console.error);
