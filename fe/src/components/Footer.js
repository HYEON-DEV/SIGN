/**
 * 푸터
 */

import styled from "styled-components";

const FooterContainer = styled.footer`
    padding: 10px;
    background-origin: var(--secondary-color);
    /* color: white; */
    text-align: center;
`;

const Footer = () => {
    return (
        <FooterContainer>
            <p> &copy; 2025 DID System </p>
        </FooterContainer>
    );
};

export default Footer;