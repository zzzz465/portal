import _ from 'lodash'
import type { NextPage } from 'next'
import Image from 'next/image'
import noImage from '../../public/no-image.png'

const Home: NextPage = () => {
  const items = (count = 32) => _.times(count).map(() =>
  (
    <li key={count}>
      <div className='flex gap-4 justify-between'>
        <h4>host1.example.com</h4>
        <ul className='flex gap-2'>
          <li>Key=Value</li>
          <li>Key2=Value2</li>
        </ul>
      </div>
    </li>
  ))

  const gridItems = (count = 32) => _.times(count).map(() => (
    <li key={count}>
      <a className='block'>
        {/* Icon */}
        <Image width='128px' height='128px' className='aspect-square' alt='icon' src={noImage.src} />

        {/* title */}
        {/* <Text> creates hydration error. https://stackoverflow.com/questions/71706064/react-18-hydration-failed-because-the-initial-ui-does-not-match-what-was-render */}
        {/* <Text>Foobar</Text> */}
        <h3 className='text-center'>NAME HERE</h3>
      </a>
    </li>
  ))

  return (
    // root
    <div style={{  /* backgroundImage: `url(${backgroundImage.src})` */ }} className='w-screen h-screen flex flex-col items-center overflow-hidden'>
      <header style={{ backgroundColor: 'green' }} className='w-full flex flex-col items-center h-14'>
        <div style={{ width: '960px', border: '2px solid red', background: 'white' }} className='h-full flex flex-row items-center my-2'>
          <h1 className='w-full text-center'>
            Search
          </h1>
        </div>
      </header>

      {/* main */}
      <div id='main-container' className='flex flex-row flex-1 overflow-auto my-12 align-middle justify-center'>
        {/* upper layout */}
        {/* <div id='upper-layout' style={{}} className='w-full h-2/5 flex flex-row'> */}
        {/* frequently accessed records */}
        {/* TODO: prevent growing */}
        {/* <div id='frequently-accessed-records' style={{ flex: '1' }} className='p-2 px-4 overflow-y-scroll'>
            <ol>
              {items()}
            </ol>
          </div> */}

        {/* recntly added records */}
        {/* <div id='recently-added-records' style={{ flex: '1' }} className='p-2 px-4 overflow-y-scroll'>
            <ol>
              {items()}
            </ol>
          </div> */}
        {/* </div> */}

        {/* bottom layout */}
        <div id='bottom-layout' style={{ flex: 1 }} className='w-3/4'>
          <div id='grid-items-container' className='mx-auto p-8'>
            {/* all groups (grid icon list) */}
            <ol className='w-full flex flex-row flex-wrap gap-12 overflow-y-hidden justify-start'>
              {gridItems()}
            </ol>
          </div>
        </div>
      </div>
    </div >
  )
}

export default Home
