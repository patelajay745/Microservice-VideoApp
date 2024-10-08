import { v2 as cloudinary } from "cloudinary";
import fs from "fs";

console.log(process.env.CLOUDINARY_CLOUD_NAME);
console.log(process.env.CLOUDINARY_ACCESS_KEY);
console.log(process.env.CLOUDINARY_ACCESS_TOKEN);
cloudinary.config({
  cloud_name: process.env.CLOUDINARY_CLOUD_NAME,
  api_key: process.env.CLOUDINARY_ACCESS_KEY,
  api_secret: process.env.CLOUDINARY_ACCESS_TOKEN,
});

export const deleteFilrOnCloudinary = async (cloudinaryImageUrl, fileType) => {
  try {
    if (!cloudinaryImageUrl) return null;
    const publicId = cloudinaryImageUrl.split("/").pop().split(".")[0];
    const response = await cloudinary.uploader.destroy(publicId, {
      resource_type: fileType,
    });

    return response;
  } catch (error) {
    console.log("File is not deleted. ", error);
  }
};

export const uploadOnCloudinary = async (localFilePath) => {
  try {
    if (!localFilePath) return null;
    const response = await cloudinary.uploader.upload(localFilePath, {
      resource_type: "auto",
    });

    fs.unlink(localFilePath, () => {});

    return response;
  } catch (error) {
    console.log("Error while uploading image on cloudinary", error);
    await fs.unlink(localFilePath, () => {});
  }
};
