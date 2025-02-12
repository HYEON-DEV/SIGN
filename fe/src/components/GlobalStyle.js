/**
 * 전역 스타일
 */

import { createGlobalStyle } from 'styled-components';

const GlobalStyle = createGlobalStyle`

  :root {
    --primary-color: #5F6F52;
    --secondary-color: #A9B388;
    --background-color: #FEFAE0;
    --accent-color: #B99470;
    --text-color: #333;
  }

  body {
    margin: 0;
    padding: 0;
    background-color: var(--background-color);
    font-family: 'Arial', sans-serif;
    color: var(--text-color);
  }

`;

export default GlobalStyle;