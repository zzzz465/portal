import { Icon, Stack, Text } from '@fluentui/react'
import type { NextPage } from 'next'
import _ from 'lodash'

const Home: NextPage = () => {
  const items = (count = 32) => _.times(count).map(() =>
  (
    <li>
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
    <li>
      <a style={{ width: '128px', height: '96px' }} className='w-full h-full block'>
        {/* Icon */}
        <image className='aspect-square w-32 h-32' />

        {/* title */}
        {/* <Text> creates hydration error. https://stackoverflow.com/questions/71706064/react-18-hydration-failed-because-the-initial-ui-does-not-match-what-was-render */}
        {/* <Text>Foobar</Text> */}
        <h3>foobar</h3>
      </a>
    </li>
  ))

  return (
    // root
    <div style={{ border: '1px solid red' }} className='w-screen h-screen flex flex-col items-center overflow-hidden'>
      <header style={{ width: '100%', height: '96px', backgroundColor: 'green', display: 'flex', flexDirection: 'column', alignItems: 'center' }}>
        <div style={{ border: '2px solid yellow', width: '960px' }} className='h-full'>
          <h1 style={{ border: '3px solid purple' }} className='w-full text-center'>
            여기에 검색창 넣기
          </h1>
        </div>
      </header>

      {/* main */}
      <div id='main-container' style={{ border: '2px solid orange', width: '960px' }} className='flex flex-col flex-1 overflow-hidden'>
        {/* upper layout */}
        <div id='upper-layout' style={{ border: '2px solid black' }} className='w-full h-2/5 flex flex-row'>
          {/* frequently accessed records */}
          {/* TODO: prevent growing */}
          <div id='frequently-accessed-records' style={{ border: '2px solid blue', flex: '1' }} className='p-2 px-4 overflow-y-scroll'>
            <ol>
              {items()}
            </ol>
          </div>

          {/* recntly added records */}
          <div id='recently-added-records' style={{ border: '2px solid red', flex: '1' }} className='p-2 px-4 overflow-y-scroll'>
            <ol>
              {items()}
            </ol>
          </div>
        </div>

        {/* bottom layout */}
        <div id='bottom-layout' style={{ border: '2px solid red', flex: 1 }} className='w-full overflow-auto'>
          {/* all groups (grid icon list) */}
          <ol className='w-full flex flex-row flex-wrap gap-8 overflow-y-hidden'>
            {gridItems()}
          </ol>
        </div>
      </div>
    </div>
  )
}

export default Home
