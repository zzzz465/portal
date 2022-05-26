import type { AppProps } from 'next/app'
import dynamic from 'next/dynamic'
import '../styles/globals.css'

const ClientInit = dynamic(() => import('../components/clientInit'), { ssr: false })

export default function MyApp({ Component, pageProps }: AppProps) {
  return (
    <>
      <ClientInit />
      <Component {...pageProps} />
    </>
  )
}
