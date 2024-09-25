import mongoose from "mongoose";

export const connectDB = async () => {
  try {
    const connectionInstance = await mongoose.connect(
      `${process.env.MONGODB_URI}${process.env.DB_NAME}`
    );

    console.log("mongoDb connected", connectionInstance.connection.name);
  } catch (error) {
    console.log("MongoDb Connection Error:", error);
  }
};
