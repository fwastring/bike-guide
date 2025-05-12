<template>
  <div class="absolute left-0 right-0 top-0 h-full z-0 pointer-events-none px-0">
    <svg class="w-full h-full" viewBox="0 220 1440 80">
      <!-- Layer 1: Main wind -->
      <path
        d="M0,240 C480,270 960,230 1440,250"
        fill="none"
        stroke="#a0aec0"
        stroke-width="8"
        opacity="0.4"
      >
        <animate
          attributeName="d"
          values="
            M0,240 C480,270 960,230 1440,250;
            M0,250 C480,260 960,240 1440,260;
            M0,240 C480,270 960,230 1440,250
          "
          dur="4s"
          repeatCount="indefinite"
        />
      </path>
      <!-- Layer 2: Subtle, thinner, faster -->
      <path
        d="M0,260 C360,280 1080,250 1440,270"
        fill="none"
        stroke="#cbd5e1"
        stroke-width="4"
        opacity="0.3"
      >
        <animate
          attributeName="d"
          values="
            M0,260 C360,280 1080,250 1440,270;
            M0,270 C360,270 1080,260 1440,280;
            M0,260 C360,280 1080,250 1440,270
          "
          dur="2.5s"
          repeatCount="indefinite"
        />
      </path>
      <!-- Layer 3: Wavy, slow, thick, more transparent -->
      <path
        d="M0,280 C600,290 840,270 1440,290"
        fill="none"
        stroke="#94a3b8"
        stroke-width="12"
        opacity="0.15"
      >
        <animate
          attributeName="d"
          values="
            M0,280 C600,290 840,270 1440,290;
            M0,290 C600,280 840,280 1440,300;
            M0,280 C600,290 840,270 1440,290
          "
          dur="6s"
          repeatCount="indefinite"
        />
      </path>
      <!-- Layer 4: Small, quick gust -->
      <path
        d="M0,300 C700,310 900,290 1440,310"
        fill="none"
        stroke="#e0e7ef"
        stroke-width="3"
        opacity="0.2"
      >
        <animate
          attributeName="d"
          values="
            M0,300 C700,310 900,290 1440,310;
            M0,310 C700,300 900,300 1440,320;
            M0,300 C700,310 900,290 1440,310
          "
          dur="1.8s"
          repeatCount="indefinite"
        />
      </path>
    </svg>
    <!-- Left gradient -->
    <div class="absolute left-0 top-0 h-full w-1/2 bg-gradient-to-r from-[#FBFBFB]/100 to-[#FBFBFB]/0 pointer-events-none"></div>
    <!-- Right gradient -->
    <div class="absolute right-0 top-0 h-full w-1/2 bg-gradient-to-l from-[#FBFBFB]/100 to-[#FBFBFB]/0 pointer-events-none"></div>
  </div>
</template> 