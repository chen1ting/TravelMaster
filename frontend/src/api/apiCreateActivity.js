import { ENDPOINT } from "./api";

const sendCreateActivityReq = async (
  descriptionlocation,
  addressactivity,
  descriptionActivity,
  activityname,
  ispaid,
  image,
  sundayopenhr,
  sundayclosehr,
  mondayopenhr,
  mondayclosehr,
  tuesdayopenhr,
  tuesdayclosehr,
  wednesdayopenhr,
  wednesdayclosehr,
  thursdayopenhr,
  thursdayclosehr,
  fridayopenhr,
  fridayclosehr,
  saturdayopenhr,
  saturdayclosehr
) => {
  console.log("?");
  const rawResponse = await fetch(ENDPOINT + "/createreviews", {
    method: "POST",
    headers: { "Content-Type": "application/json" },
    body: JSON.stringify({
      categories: descriptionlocation,
      latitude: addressactivity,
      description: descriptionActivity,
      title: activityname,
      paid: ispaid,
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

export { sendCreateActivityReq };
