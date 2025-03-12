import { useCookies } from '@vueuse/integrations/useCookies'

const authCookie = useCookies(['auth'])

export function setAuthCookie(name: string) {
    authCookie.set('auth', name)
}

export function getAuthCookie() {
    return authCookie.get('auth')
}

export function removeAuthCookie() {
    return authCookie.remove('auth')
}