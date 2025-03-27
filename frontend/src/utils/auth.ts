import { useCookies } from '@vueuse/integrations/useCookies'

const authCookie = useCookies(['auth'])

type User = {
    name: string
    role: string
    ID: number
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

export function getUserName() {
    return authCookie.get('auth').name
}

export function getAuthCookie() {
    return authCookie.get('auth')
}

export function removeAuthCookie() {
    return authCookie.remove('auth')
}