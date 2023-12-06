export const capitalize = (str: string) => {
    return str.charAt(0).toUpperCase() + str.slice(1);
}

export const extractHostname = (url: string): string => {
    try {
        const parsedUrl: URL = new URL(url);
        let hostname: string = parsedUrl.hostname;
        const parts: string[] = hostname.split('.');
        return parts.slice(0, -1).join('.');
    } catch (error) {
        console.error("Invalid URL:", error);
        return "";
    }
}