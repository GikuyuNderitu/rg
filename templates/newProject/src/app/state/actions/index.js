export function getDiscussions() {
    return new Promise((resolve, reject) => {
        setTimeout(()=>resolve([
            {
                summary: "firstOne",
                date: "5 days ago"
            },
            {
                summary: "second One boi",
                date: "Yesterday"
            }
        ]), 1000);
    })
}