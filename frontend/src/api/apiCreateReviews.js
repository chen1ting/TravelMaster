const ENDPOINT = "http://localhost:8080";
const sendCreateReviewReq = async (name, date, profileImg, reviewTitle, reviewBody, uploadImg1, uploadImg2, uploadImg3) => {
    console.log("?");
    const rawResponse = await fetch(ENDPOINT + "/createreviews", {
        method: "POST",
        headers: { "Content-Type": "application/json" },
        body: JSON.stringify({
            username: name,
            date: date,
            profileImg: profileImg,
            reviewTitle: reviewTitle,
            reviewBody: reviewBody,
            uploadImg1: uploadImg1,
            uploadImg2: uploadImg2,
            uploadImg3: uploadImg3
        }),
    });

    if (rawResponse.status !== 201) {
        console.log("resp: " + rawResponse.status); // TODO: might wanna return an err message to display here
        return null;
    }
    const content = await rawResponse.json();
    return content;
};

export { sendCreateReviewReq };