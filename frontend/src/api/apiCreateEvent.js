const ENDPOINT = "http://localhost:8080";
const sendCreateEventReq = async (descriptionlocation, addressevent, descriptionevent, eventname, sundayopenhr
    , sundayclosehr, mondayopenhr, mondayclosehr, tuesdayopenhr, tuesdayclosehr, wednesdayopenhr, wednesdayclosehr
    , thursdayopenhr, thursdayclosehr, fridayopenhr, fridayclosehr, saturdayopenhr, saturdayclosehr) => {
    console.log("?");
    const rawResponse = await fetch(ENDPOINT + "/createreviews", {
        method: "POST",
        headers: { "Content-Type": "application/json" },
        body: JSON.stringify({
            category: descriptionlocation,
            latitude: addressevent,
            description: descriptionevent,
            title: eventname,
            sun_opening_time: sundayopenhr,
            sun_closing_time: sundayclosehr,
            mon_opening_time: mondayopenhr,
            mon_closing_time: mondayclosehr,
            tue_opening_time: tuesdayopenhr,
            tue_closing_time: tuesdayclosehr,
            wed_opening_time: wednesdayopenhr,
            wed_closing_time: wednesdayclosehr,
            thur_opening_time: thursdayopenhr,
            thur_closing_time: thursdayclosehr,
            fri_opening_time: fridayopenhr,
            fri_closing_time: fridayclosehr,
            sat_opening_time: saturdayopenhr,
            sat_closing_time: saturdayclosehr,
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