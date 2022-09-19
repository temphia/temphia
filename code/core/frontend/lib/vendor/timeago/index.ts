// vendored => https://github.com/mikepenzin/timeago-simple/blob/master/index.js
export const time_ago = (date_time) => {
    // today date and time in milliseconds 
    var today = Date.now();
    var dateParse = Date.parse(date_time);

    //We will perform some test - if there is error, we will throw error to console and exit, no change will be on the data.
    try {
        // We need to check if we able to parse the Date (if the result is NaN, this is an issue)
        if (dateParse !== dateParse) throw "timeago-simple: Please check date and time format! Unable to parse the date & time: " + date_time;
    }
    catch (err) {
        console.error(err);
        return (date_time);
    }

    if ((dateParse - today) < 0) {
        return pastCalc(date_time);
    } else {
        return futureCalc(date_time);
    }
};


// General help functions for time calculations
function pastCalc(timeData) {

    // today date and time in milliseconds 
    var today = Date.now();

    // parsing post date and time into milliseconds format
    timeData = Date.parse(timeData);

    var seconds = (today - timeData) / 1000;
    var minutes = (seconds / 60);
    var hours = (seconds / 3600);
    if (seconds < 60 && minutes < 1) {
        return (seconds === 1 ? Math.round(seconds) + " second ago" : Math.round(seconds) + " seconds ago");
    }
    if (minutes < 60 && hours < 1) {
        return (minutes === 1 ? Math.round(minutes) + " minute ago" : Math.round(minutes) + " minutes ago");
    }
    if (hours > 24) {
        var days = hours / 24;
        if (days > 30) {
            var month = days / 30;
            if (month > 12) {
                var years = month / 12;
                if (years > 0) {
                    return (years === 1 ? Math.ceil(years) + " year ago" : Math.ceil(years) + " years ago");
                }
            }
            return (Math.round(month) + " month ago");
        }
        return (days === 1 ? Math.round(days) + " day ago" : Math.round(days) + " days ago");
    } else {
        return (hours === 1 ? Math.round(hours) + " hour ago" : Math.round(hours) + " hours ago");
    }

}

function futureCalc(timeData) {

    // today date and time in milliseconds 
    var today = Date.now();

    // parsing post date and time into milliseconds format
    timeData = Date.parse(timeData);
    var seconds = (timeData - today) / 1000;
    var minutes = (seconds / 60);
    var hours = (seconds / 3600);
    if (seconds < 60 && minutes < 1) {
        return (seconds === 1 ? "in " + Math.round(seconds) + " second" : "in " + Math.round(seconds) + " seconds");
    }
    if (minutes < 60 && hours < 1) {
        return (minutes === 1 ? "in " + Math.round(minutes) + " minute" : "in " + Math.round(minutes) + " minutes");
    }
    if (hours > 24) {
        var days = hours / 24;
        if (days > 30) {
            var month = days / 30;
            if (month > 12) {
                var years = month / 12;
                if (years > 0) {
                    return (years === 1 ? "in " + Math.ceil(years) + " year" : "in " + Math.ceil(years) + " years");
                }
            }
            return ("in " + Math.round(month) + " month");
        }
        return (days === 1 ? "in " + Math.round(days) + " day" : "in " + Math.round(days) + " days");
    } else {
        return (hours === 1 ? "in " + Math.round(hours) + " hour" : "in " + Math.round(hours) + " hours");
    }
}

