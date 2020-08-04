import Service from '../util/axios';

function GetNativeTokenName(){
    const url = '/api/basedenom';
    return Service.http(url);
}

export {
    GetNativeTokenName
}