export const handler = async(event) => {
    // TODO implement
    const response = {
        statusCode: 200,
        body: JSON.stringify('Hello from NodeJS Handler in GO-CDK'),
    };
    return response;
};