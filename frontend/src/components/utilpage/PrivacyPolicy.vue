<template>
  <div style="width:100%;">
    <div class="privacy_policy_wrap">
      <div class="privacy_policy_wrap_title">
        <h2 :style="devicesWidth<=500?'width:100%;':''">Privacy Policy</h2>
      </div>
      <div :class="privacyPolicyVar">
        <h4 style="margin-top:0.4rem;">PLEASE READ THE PRIVACY POLICY CAREFULLY. (July 2018)</h4>
        <p>This Privacy Policy is designed to assist you with understanding how IRIS Foundation Limited (referred to as "the Foundation", "we", "us" or "our") collect, use and disclose the information you provide us when accessing or using the website,<a href="">irisplorer.io</a>.</p>
        <h4>WHAT INFORMATION DO WE COLLECT?</h4>
        <p>You may provide us your personal information that you voluntarily disclose that is collected on an individual basis; and information collected automatically when you use our Website or the services available on our Website. Each time you use any Site, or any Services, the current version of this Privacy Policy will apply. Personal Information may include (but is not limited to):</p>
        <ul>
          <li>personal information that links back to an individual, e.g., name, E-mail address, cryptocurrency address, and other personal identification information;</li>
          <li>technical information, e.g., IP address for API services and login;</li>
          <li>statistical data, e.g., hits to website.</li>
        </ul>
        <h4>HOW DO WE COLLECT PERSONAL INFORMATION?</h4>
        <p>From our website, we automatically collect information from your devices. We collect the following information:</p>
        <ul>
          <li>IP address</li>
        </ul>
        <p>We use your IP address to help diagnose problems with our server, and to administer our website.</p>
        <ul>
          <li>Cookies</li>
        </ul>
        <p>A cookie is an element of data that a website can send to your browser, which may then store it on your system. We use cookies in some of our pages to store your preferences and record session information. The information that we collect is then used to ensure a more personalized service level for our users.</p>
        <p>You can adjust settings on your browser so that you will be notified when you receive a cookie. Please refer to your browser documentation to check if cookies have been enabled on your computer or to request not to receive cookies. As cookies allow you to take advantage of some of the Website’s essential features, we recommend that you accept cookies.</p>
        <ul>
          <li>User feedback form</li>
        </ul>
        <p>Our feedback form requires you to give us contact information (e.g. your name and email address) so that we can respond to your comments. We use your contact information from the registration form to send you information about our company. Your contact information is also used to contact you where necessary.</p>
        <h4>USE OF INFORMATION</h4>
        <p>We may use your Personal Information for the following purposes:</p>
        <ul>
          <li>providing, maintaining, delivering or improving our Website or the products or services provided through our Website;</li>
          <li>analyzing and tracking data to determine the usefulness or popularity of certain content and to better understand the online activity of our Website users;</li>
          <li>investigate and respond to claims and inquiries from you;</li>
          <li>providing you with the information or products or services that you have requested;</li>
          <li>answering your inquiry or responding to a communication from you;</li>
          <li>developing new products or services;</li>
          <li>sending you technical notices, support or administrative notifications;</li>
          <li>communicating with you about news, products, services, events and other information we think will be of interest to you;</li>
          <li>monitoring and analyzing trends, usage and activities in connection with our Website;</li>
          <li>detecting, investigating and preventing fraudulent transactions or unauthorized or illegal activities;</li>
          <li>protecting the rights and property of the Foundation Parties (as defined below) and others;</li>
          <li>for all other purposes ancillary to any of the purposes stated above.</li>
        </ul>
        <h4>INFORMATION DISCLOSING</h4>
        <p>We will not trade or sell your Personal Information to third parties. Your Personal Information shall only be disclosed or transferred in the following ways:</p>
        <ul>
          <li>with your consent or at your instruction;</li>
          <li>with our current or future parent companies, affiliates, subsidiaries and with other companies under common control or ownership with us or our offices internationally;</li>
          <li>in connection with a merger or sale of our company assets, or if we do a financing or are involved in an acquisition or any other situation where Information may be disclosed or transferred as one of our business assets;</li>
          <li>in response to a request for information if we believe disclosure is in accordance with, or required by, any applicable law, regulation or legal process;</li>
          <li>if we believe your actions are inconsistent with our user agreements or policies, or to protect the rights, property and safety of any block.one Party or others;</li>
          <li>with third parties where aggregate Information is disclosed which cannot be reasonably be used to identify you.</li>
        </ul>
        <p>We shall take practical steps to ensure that their employees, officers, agents, consultants, contractors and such other third parties mentioned above who are involved in the collection, use and disclosure of your Personal Data will observe and adhere to the terms of this Privacy Statement. </p>
        <h4>PROTECTION OF PERSONAL INFORMATION</h4>
        <p>We take reasonable steps to protect the security of the information communicated through our Website. However, no computer security system is entirely foolproof and the Internet is not a secure method of transmitting information. As a result, we do not assume any responsibility for the data you submit to or receive from us through the Internet or for any unauthorised access or use of that information and we cannot and do not guarantee that information communicated by you to us or sent to you by us will be received or that it will not be altered before or after its transmission to us. You agree to not hold us and our respective past, present and future employees, officers, directors, contractors, consultants, equity holders, suppliers, vendors, service providers, parent companies, subsidiaries, affiliates, agents, representatives, predecessors, successors and assigns (collectively, the “block.one Parties”) liable for any loss or damage of any sort incurred as a result of any misappropriation, interception, modification, deletion, destruction or use of information provided through our Website.</p>

        <h4>HYPERLINKS AND THIRD-PARTY SITES</h4>
        <p>This Website may contain links to other third party websites. This Privacy Notice does not apply to such Third Party Sites as those sites are outside our control. If you access Third Party Sites using the links provided, the operators of these sites may collect your personal information, including through cookies or other technologies. You should consult these other websites' privacy policies as we have no control over them and are not responsible for any information that is submitted to or collected by these third parties.</p>

        <h4>ADDITIONAL INFORMATION OR ASSISTANCE</h4>
        <p>We may change this Privacy Policy in our sole discretion at any time. Any revisions to this Privacy Policy will be posted on the home page of our Website. It is your obligation to periodically visit our Website to review any changes that may be made to this Privacy Policy. Your continued use of our Website constitutes your agreement to be bound by any such changes to this Privacy Policy. If you do not agree with this Privacy Policy, please do not use the Wesite, or any of our Services.</p>
        <p>If you have any questions about this Privacy Policy, please contact us at contact@irisnet.org.</p>



      </div>
    </div>

  </div>
</template>

<script>
  import Tools from "../../util/Tools";

  export default {
    data() {
      return {
        devicesWidth: window.innerWidth,
        privacyPolicyVar: 'person_computer_privacy_policy_wrap',

      }
    },
    watch:{
      devicesWidth(devicesWidth){
        if (devicesWidth > 500) {
          this.privacyPolicyVar = 'person_computer_privacy_policy_wrap';
        } else {
          this.privacyPolicyVar = 'mobile_privacy_policy_wrap';
        }
      }
    },
    beforeMount() {
      Tools.scrollToTop();
      if (this.devicesWidth > 500) {
        this.privacyPolicyVar = 'person_computer_privacy_policy_wrap';
      } else {
        this.privacyPolicyVar = 'mobile_privacy_policy_wrap';
      }
    },
  }
</script>
<style lang="scss">
  @import '../../style/mixin';
  .privacy_policy_wrap{
    width:100%;
    @include flex;
    flex-direction:column;
    align-items: center;

    .privacy_policy_wrap_title{
      @include flex;
      flex-direction:column;
      align-items:center;
      width:100%;
      padding-left:0.2rem;
      background: #F5F7FD;
      h2{
        font-size:0.22rem;
        margin-bottom:0;
        height:0.62rem;
        line-height:0.62rem;

        width:80%;
      }
    }
    .person_computer_privacy_policy_wrap{
      padding-left:0.2rem;
      width:80%;
      margin-bottom:0.4rem;
    }
    .mobile_privacy_policy_wrap{
      width:100%;
      padding:0.15rem;
    }
    a{
      color: var(--bgColor)!important;
    }

    h4{
      font-size:0.18rem;
      margin-top:0.2rem;
      line-height:0.22rem;
      margin-bottom:0;

    }
    ul{
      margin-left:0.2rem;
    }
    p, li{
      list-style: inherit!important;
      font-size:0.14rem;
      color:var(--contentColor);
    }
  }


</style>
