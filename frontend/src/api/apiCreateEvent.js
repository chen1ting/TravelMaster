const ENDPOINT = "http://localhost:8080";
const sendCreateEventReq = async (descriptionlocation, addressevent, descriptionevent, eventname, sundayopenhr
    , sundayclosehr, mondayopenhr, mondayclosehr, tuesdayopenhr, tuesdayclosehr, wednesdayopenhr, wednesdayclosehr
    , thursdayopenhr, thursdayclosehr, fridayopenhr, fridayclosehr, saturdayopenhr, saturdayclosehr) => {
    console.log("?");
    const rawResponse = await fetch(ENDPOINT + "/createreviews", {
        method: "POST",
        headers: { "Content-Type": "application/json" },
        body: JSON.stringify({
            descriptionlocation: descriptionlocation,
            addressevent: addressevent,
            descriptionevent: descriptionevent,
            eventname: eventname,
            sundayopenhr: sundayopenhr,
            sundayclosehr: sundayclosehr,
            mondayopenhr: mondayopenhr,
            mondayclosehr: mondayclosehr,
            tuesdayopenhr: tuesdayopenhr,
            tuesdayclosehr: tuesdayclosehr,
            wednesdayopenhr: wednesdayopenhr,
            wednesdayclosehr: wednesdayclosehr,
            thursdayopenhr: thursdayopenhr,
            thursdayclosehr: thursdayclosehr,
            fridayopenhr: fridayopenhr,
            fridayclosehr: fridayclosehr,
            saturdayopenhr: saturdayopenhr,
            saturdayclosehr: saturdayclosehr,
        }),
    });

    if (rawResponse.status !== 201) {
        console.log("resp: " + rawResponse.status); // TODO: might wanna return an err message to display here
        return null;
    }
    const content = await rawResponse.json();
    return content;
};

export { sendCreateEventReq };