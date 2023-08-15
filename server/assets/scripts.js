const relativeTimeFormat = new Intl.RelativeTimeFormat();

const formatRelativeHours = (timestamp) => {
  const seconds = timestamp - Date.now() / 1000;
  const hours = seconds / 60 / 60;
  return relativeTimeFormat.format(Math.round(hours), "hours");
};

htmx.onLoad(function (content) {
  const timestamps = content.querySelectorAll("[data-timestamp]");

  for (const timestamp of timestamps) {
    timestamp.textContent = formatRelativeHours(
      parseInt(timestamp.textContent),
    );
  }
});
