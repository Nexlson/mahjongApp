import Head from 'next/head'
import Footer from '../components/Footer'
import NavBar from '../components/Navbar'

export default function Home() {
  return (
    <div>
      <Head>
        <title>Mahjong App</title>
        <meta name="description" content="Simple mahjong app" />
        <link rel="icon" href="/mahjong.ico" />
      </Head>
      <NavBar />
      <h1>This is some contents for main page</h1>
      <Footer/>
    </div>
  )
}
