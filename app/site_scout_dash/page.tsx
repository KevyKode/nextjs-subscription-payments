"use client";

import CustomerPortalForm from '@/components/ui/AccountForms/CustomerPortalForm';
import EmailForm from '@/components/ui/AccountForms/EmailForm';
import NameForm from '@/components/ui/AccountForms/NameForm';
import { useEffect, useState } from 'react';
import { powerbiApi } from '@/powerbi.api'; // Adjust the import path as necessary

export default function SiteScoutDash() {
  const [dashboardUrl, setDashboardUrl] = useState<string | null>(null); // Correctly typed useState
  const [userDetails, setUserDetails] = useState<any>(null);
  const [subscription, setSubscription] = useState<any>(null);
  const [user, setUser] = useState<any>(null);

  useEffect(() => {
    async function fetchDashboard() {
      try {
        const res = await fetch('/api/webhooks/supabase');
        if (!res.ok) {
          throw new Error(`HTTP error! status: ${res.status}`);
        }
        const data = await res.json();

        console.log('Fetched Data:', data);
        setUser(data.user);
        setUserDetails(data.userDetails);
        setSubscription(data.subscription);

        // Fetch Power BI Access Token
        const accessToken = await powerbiApi.authenticate();
        console.log('Access Token:', accessToken);
        
        // Set the dashboard URL with the access token
        const embedUrl = `https://app.powerbi.com/reportEmbed?groupId=90eb8937-dc8b-409a-877b-fc25155669b0&dashboardId=3593492a-f7ea-49cc-b736-7da4829364e8&access_token=${accessToken}`;
        setDashboardUrl(embedUrl);
        console.log('Dashboard URL:', embedUrl);
      } catch (error) {
        console.error('Fetch Error:', error);
      }
    }

    fetchDashboard();
  }, []);

  return (
    <section className="mb-32 bg-black">
      <div className="max-w-6xl px-4 py-8 mx-auto sm:px-6 sm:pt-24 lg:px-8">
        <div className="sm:align-center sm:flex sm:flex-col">
          <h1 className="text-4xl font-extrabold text-white sm:text-center sm:text-6xl">
            Site Scout Dashboard
          </h1>
          <p className="max-w-2xl m-auto mt-5 text-xl text-zinc-200 sm:text-center sm:text-2xl">
            We partnered with Stripe for a simplified billing.
          </p>
        </div>
      </div>
      <div className="p-4">
        {subscription && <CustomerPortalForm subscription={subscription} />}
        {userDetails && <NameForm userName={userDetails.full_name ?? ''} />}
        {user && <EmailForm userEmail={user.email} />}
        {dashboardUrl && (
          <iframe
            src={dashboardUrl}
            frameBorder="0"
            width="100%"
            height="800"
            allowFullScreen={true}
          />
        )}
      </div>
    </section>
  );
}