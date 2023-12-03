declare module "*.module.css";

declare module "*.svg" {
    const content: any;
    export default content;
}
interface StringMap {
    [key: string]: string;
}