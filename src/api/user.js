import request from "@/utils/request"

export function login(data) {
    return request({
        url: "/user/login",
        method: 'post',
        data
    })
}

export function logout() {
    return request({
        url: "/user/logout",
        method: 'get',
    })
}

export function myInfo() {
    return request({
        url: "/user/my",
        method: 'get',
    })
} 