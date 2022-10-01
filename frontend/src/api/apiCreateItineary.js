const ENDPOINT = "http://localhost:8080";
const sendCreateItinearyReq = async (descriptionlocation, addressevent, descriptionevent, eventname, visitdatestart, visitdateend) => {
  console.log("?");
  const rawResponse = await fetch(ENDPOINT + "/createitineary", {
    method: "POST",
    headers: { "Content-Type": "application/json" },
    body: JSON.stringify({
      descriptionlocation: descriptionlocation,
      addressevent: addressevent,
      descriptionevent: descriptionevent,
      eventname: eventname,
      visitdatestart: visitdatestart,
      visitdateend: visitdateend,
    }),
  });

    if (rawResponse.status !== 201) {
        console.log("resp: " + rawResponse.status); // TODO: might wanna return an err message to display here
        return null;
    }
    const content = await rawResponse.json();
    return content;
};
    export { sendCreateItinearyReq };