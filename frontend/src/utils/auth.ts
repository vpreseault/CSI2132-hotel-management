import { useCookies } from '@vueuse/integrations/useCookies'

const authCookie = useCookies(['auth'])

type User = {
    name: string
    role: string
    ID: number
    address: string
}

export function setAuthCookie(user: User) {
    authCookie.set('auth', user)
}

export function getUserRole() {
    return authCookie.get('auth').role
}

export function getUserID() {
    return authCookie.get('auth').ID
}

export function getUserAddress() {
    return authCookie.get('auth').address
}

export function getUserName() {
    return authCookie.get('auth').name
}

export function getAuthCookie() {
    return authCookie.get('auth')
}

export function removeAuthCookie() {
    return authCookie.remove('auth')
}